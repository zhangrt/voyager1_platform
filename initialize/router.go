package initialize

import (
	"net/http"
	"time"

	auth "github.com/zhangrt/voyager1_core/auth/luna"
	_ "github.com/zhangrt/voyager1_platform/docs"
	"github.com/zhangrt/voyager1_platform/global"
	middleware "github.com/zhangrt/voyager1_platform/middleware"
	"github.com/zhangrt/voyager1_platform/router"
	service "github.com/zhangrt/voyager1_platform/service/auth"

	handler "github.com/zhangrt/voyager1_core/auth/luna/handler"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {

	switch global.GS_CONFIG.System.Mode {
	case "develop":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	Router := gin.Default()
	testRouter := router.RouterGroupApp.Test
	fileRouter := router.RouterGroupApp.File
	systemRouter := router.RouterGroupApp.System
	demoRouter := router.RouterGroupApp.Demo
	statisticsRouter := router.RouterGroupApp.Statistics

	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors())        // 直接放行全部跨域请求
	Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.GS_LOG.Info("Cors init")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GS_LOG.Info("register swagger handler")

	// 重定向到swagger
	Router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, global.GS_CONFIG.System.Application)
	}).GET(global.GS_CONFIG.System.Application, func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})

	PublicGroup := Router.Group(global.GS_CONFIG.System.Application)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "Okay : "+time.Stamp)
		})
	}
	{
		testRouter.InitTestRouter(PublicGroup)     // 测试路由
		demoRouter.InitFacilityRouter(PublicGroup) // 演示demo测试路由
		systemRouter.InitBaseRouter(PublicGroup)   // 注册登录基础路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup)   // 初始化相关路由

	}

	PrivateGroup := Router.Group(global.GS_CONFIG.System.Application)
	// 权限管理 test模式下跳过
	if global.GS_CONFIG.System.Mode != "test" {
		// 注册权限管理模块，注入实现类
		auth.RegisterCasbin(&service.CasbinService{})                    // 注入Casbin实现类
		auth.RegisterJwt(&service.JwtService{})                          // 注入Jwt实现类
		auth.NewJWT().LoadAll()                                          // 加载黑名单
		PrivateGroup.Use(handler.JWTAuth()).Use(handler.CasbinHandler()) // 注入拦截器
	}

	{
		fileRouter.InitFileRouter(PrivateGroup)                  // 文件上传下载相关路由
		systemRouter.InitSystemRouter(PrivateGroup)              // system相关路由
		systemRouter.InitUserRouter(PrivateGroup)                // 用户相关路由
		systemRouter.InitMenuRouter(PrivateGroup)                // 菜单相关路由
		systemRouter.InitJwtRouter(PrivateGroup)                 // jwt相关路由
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
		systemRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理相关路由
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色相关路由
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)  // 注册角色按钮相关路由
		systemRouter.InitWeatherrRouter(PrivateGroup)            // 天气信息相关路由
		statisticsRouter.InitStatisticesRouter(PrivateGroup)     // 统计数据相关路由
	}

	InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	global.GS_LOG.Info("router register success")
	return Router
}
