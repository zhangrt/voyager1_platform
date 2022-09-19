package demo

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type FacilityRouter struct{}

func (e *FacilityRouter) InitFacilityRouter(Router *gin.RouterGroup) {
	router := Router.Group("facility")
	facilityApi := v1.ApiGroupApp.DemoApiGroup.FacilityApi
	{
		router.POST("add", facilityApi.AddFacility)         // 添加设备
		router.DELETE("remove", facilityApi.RemoveFacility) // 删除设备
		router.PUT("update", facilityApi.UpdateFacility)    // 修改设备
		router.GET("get", facilityApi.GetFacility)          // 获取设备详情
		router.GET("getList", facilityApi.GetFacilityList)  // 获取设备列表
	}
}
