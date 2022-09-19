package statistics

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type StatisticsRouter struct{}

func (e *StatisticsRouter) InitStatisticesRouter(Router *gin.RouterGroup) {
	router := Router.Group("statistics")
	statisticsApi := v1.ApiGroupApp.StatisticsApiGroup.StatisticsApi
	{
		router.GET("facilityByType", statisticsApi.StatisticsFacilityByType) // 统计不同类型设备数量
	}
}
