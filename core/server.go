package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	core "github.com/zhangrt/voyager1_core"
	"github.com/zhangrt/voyager1_core/cache"
	config "github.com/zhangrt/voyager1_core/config"
	"github.com/zhangrt/voyager1_platform/global"
	initialize "github.com/zhangrt/voyager1_platform/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {

	// 基础组件对象初始化
	core.New().
		Viper(global.GS_VP).
		Zap(global.GS_LOG).
		// DB(global.GS_DB). // core组件去除了数据库相关代码
		SetRedisMod(global.GS_CONFIG.Redis.ClusterMod).
		RedisStandalone(global.GS_REDIS_STANDALONE).
		// BlackCache(global.BlackCache).
		Config(config.Server{
			System:  config.System(global.GS_CONFIG.System),
			JWT:     global.GS_CONFIG.JWT,
			Casbin:  global.GS_CONFIG.Casbin,
			AUTHKey: global.GS_CONFIG.AUTHKey,
		}).
		ConfigMinio(global.GS_CONFIG.Minio).
		ConfigZinx(global.GS_CONFIG.Zinx).
		ConfigGrpc(global.GS_CONFIG.Grpc).
		ConfigCache(global.GS_CONFIG.Cache)

	if global.GS_CONFIG.System.UseCache {
		// 新版 cache封装初始化
		global.GS_CACHE = cache.CreateCache()
		// 旧版 redis初始化
		// if global.GS_CONFIG.System.CacheType == "redis" {
		// 	if global.GS_CONFIG.Redis.ClusterMod {
		// 		// 初始化redis集群
		// 		initialize.RedisCLuster()
		// 	} else {
		// 		// 初始化单机redis服务
		// 		initialize.Redis()
		// 	}
		// }
	}
	if global.GS_CONFIG.System.UseDatabase {
		// 初始化数据库
		initialize.Gorm()
		// 初始化自动建表
		if global.GS_CONFIG.System.AutoMigrate {
			initialize.AutoMigrate()
		}
		// 程序结束前关闭数据库链接
		db, _ := global.GS_DB.DB()
		defer db.Close()
	}

	// 启动 Luan Server(Grpc/Tcp)
	// go grpc.NewServer().
	// 	RegisterAuthServiceServer(new(service.AuthService)).
	// 	LunchGrpcServer()

	// 时区
	time.LoadLocation(global.GS_CONFIG.System.TimeZone)

	// 路由
	Router := initialize.Routers()

	Addr := fmt.Sprintf("%s:%s", global.GS_CONFIG.System.Host, global.GS_CONFIG.System.Port)
	// init
	s := func(address string, router *gin.Engine) server {
		return &http.Server{
			Addr:           Addr,
			Handler:        router,
			ReadTimeout:    20 * time.Second,
			WriteTimeout:   20 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
	}(Addr, Router)

	global.GS_LOG.Info("server run success on ", zap.String("address", Addr))

	fmt.Printf(`
                                 ^__^       /
                         _______/(oo)  ___ /
                     /\\/(       /(__)
                         | w----|\\
                        /\\     |/

	welcome to Platform
	version:v0.1
	email:zhoujiajun@gsafety.com
	default docs:http://%s/swagger/index.html

`, Addr)
	global.GS_LOG.Error(s.ListenAndServe().Error())
}
