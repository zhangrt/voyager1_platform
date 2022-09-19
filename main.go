package main

import (
	"github.com/zhangrt/voyager1_platform/global"

	"github.com/zhangrt/voyager1_platform/core"

	"go.uber.org/zap"
)

// @title Swagger github.com/zhangrt/voyager1_platform demo API
// @version 0.0.1
// @description Platform demo golang 后端服务api
// @in header
// @BasePath /github.com/zhangrt/voyager1_platform
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
func main() {

	global.GS_VP = core.Viper() // 初始化Viper
	global.GS_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GS_LOG)

	// 启动服务
	core.RunServer()
}
