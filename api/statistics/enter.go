package statistics

import "github.com/zhangrt/voyager1_platform/service"

type ApiGroup struct {
	StatisticsApi
}

var (
	statisticsFacilityService = service.ServiceGroupApp.StatisticsGroup.StatisticsFacilityService
)
