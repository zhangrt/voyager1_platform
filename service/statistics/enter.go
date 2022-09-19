package statistics

import "github.com/zhangrt/voyager1_platform/repository"

type ServiceGroup struct {
	StatisticsFacilityService
	// StatisticsAlarmService
}

var (
	statisticsFacilityRepository = repository.RepositoryGroupApp.StatisticsRepository.StatisticsFacilityRepository
	// tatisticsAlarmRepository     = repository.RepositoryGroupApp.StatisticsRepository.StatisticsAlarmRepository
)
