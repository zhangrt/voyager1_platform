package statistics

import (
	"strings"

	"github.com/zhangrt/voyager1_platform/model/statistics"
)

type StatisticsFacilityService struct{}

func (statisticsFacilityService *StatisticsFacilityService) StatisticsFacilityByType(param statistics.StatisticsFacilityByType) (s []statistics.StatisticsFacilityType, err error) {
	var results []statistics.StatisticsFacilityType

	results, err = statisticsFacilityRepository.StatisticsFacilityByType(param)

	if len(param.Types) == 0 {
		return results, err
	}

	var types = strings.Split(param.Types, ",")

	var data []statistics.StatisticsFacilityType = []statistics.StatisticsFacilityType{}

	for i := range types {

		for j := range results {
			if types[i] == results[j].Type {
				data = append(data, results[j])
			}
		}
	}
	return data, err
}
