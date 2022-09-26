package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zhangrt/voyager1_platform/global"
	initialize "github.com/zhangrt/voyager1_platform/initialize"

	"github.com/gin-gonic/gin"
	gallery "github.com/zhangrt/voyager1_core"
	auth "github.com/zhangrt/voyager1_core/auth/luna"
	config "github.com/zhangrt/voyager1_core/config"
	s "github.com/zhangrt/voyager1_core/zinx/server"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GS_CONFIG.System.UseCache {
		if global.GS_CONFIG.System.CacheType == "redis" {
			if global.GS_CONFIG.Redis.ClusterMod {
				// 初始化redis集群
				initialize.RedisCLuster()
			} else {
				// 初始化单机redis服务
				initialize.Redis()
			}
		}
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

	// 基础组件对象初始化
	gallery.NewInit().
		Viper(global.GS_VP).
		Zap(global.GS_LOG).
		DB(global.GS_DB).
		SetRedisMod(global.GS_CONFIG.Redis.ClusterMod).
		RedisStandalone(global.GS_REDIS_STANDALONE).
		BlackCache(global.BlackCache).
		Config(config.Server{
			System:  config.System(global.GS_CONFIG.System),
			JWT:     global.GS_CONFIG.JWT,
			Casbin:  global.GS_CONFIG.Casbin,
			AUTHKey: global.GS_CONFIG.AUTHKey,
		}).
		ConfigMinio(global.GS_CONFIG.Minio).
		ConfigZinx(global.GS_CONFIG.Zinx)

	auth.LoadAll()

	// 启动 Luan TCP Server
	go s.Luna()

	// 时区
	time.LoadLocation(global.GS_CONFIG.System.TimeZone)

	// 路由
	Router := initialize.Routers()

	Addr := global.GS_CONFIG.System.Host + ":" + global.GS_CONFIG.System.Port
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

	welcome to gin-github.com/zhangrt/voyager1_platform
	version:v0.1
	email:zhoujiajun@github.com/zhangrt/voyager1_platform.com
	default docs:http://%s/swagger/index.html

`, Addr)
	global.GS_LOG.Error(s.ListenAndServe().Error())
}
