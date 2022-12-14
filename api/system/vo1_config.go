package system

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/response"
	"github.com/zhangrt/voyager1_platform/model/system"
	systemRes "github.com/zhangrt/voyager1_platform/model/system/response"
	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApi struct{}

// @Tags System
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=systemRes.SysConfigResponse,msg=string} "获取配置文件内容,返回包括系统配置"
// @Router /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	if config, err := systemConfigService.GetSystemConfig(); err != nil {
		global.GS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
	}
}

// @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.Config true "设置配置文件内容"
// @Success 200 {object} response.Response{data=string} "设置配置文件内容"
// @Router /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.Config
	_ = c.ShouldBindJSON(&sys)
	if err := systemConfigService.SetSystemConfig(sys); err != nil {
		global.GS_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{msg=string} "重启系统"
// @Router /system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.GS_LOG.Error("重启系统失败!", zap.Error(err))
		response.FailWithMessage("重启系统失败", c)
	} else {
		response.OkWithMessage("重启系统成功", c)
	}
}

// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取服务器信息"
// @Router /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	if server, err := systemConfigService.GetServerInfo(); err != nil {
		global.GS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}
