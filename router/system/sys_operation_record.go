package system

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("vo1OperationRecord")
	vo1OperationRecordApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("createVo1OperationRecord", vo1OperationRecordApi.CreateVo1OperationRecord)             // 新建Vo1OperationRecord
		operationRecordRouter.DELETE("deleteVo1OperationRecord", vo1OperationRecordApi.DeleteVo1OperationRecord)           // 删除Vo1OperationRecord
		operationRecordRouter.DELETE("deleteVo1OperationRecordByIds", vo1OperationRecordApi.DeleteVo1OperationRecordByIds) // 批量删除Vo1OperationRecord
		operationRecordRouter.GET("findVo1OperationRecord", vo1OperationRecordApi.FindVo1OperationRecord)                  // 根据ID获取Vo1OperationRecord
		operationRecordRouter.GET("getVo1OperationRecordList", vo1OperationRecordApi.GetVo1OperationRecordList)            // 获取Vo1OperationRecord列表

	}
}
