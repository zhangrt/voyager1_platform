package statistics

import (
	"time"

	"github.com/zhangrt/voyager1_platform/model/common/response"
	"github.com/zhangrt/voyager1_platform/model/statistics"

	"github.com/gin-gonic/gin"
)

type StatisticsApi struct{}

// @Tags StatisticsApi
// @Summary 统计一段时间内不同类型设备数量
// @accept application/json
// @Produce application/json
// @Param data query statistics.StatisticsFacilityByType true "统计不同类型设备数量"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /statistics/facilityByType [get]
func (s *StatisticsApi) StatisticsFacilityByType(c *gin.Context) {
	var param statistics.StatisticsFacilityByType
	startTime, err1 := time.ParseInLocation("2006-01-02 15:04:05", c.Query("startTime"), time.Local)
	if err1 == nil {
		param.StartTime = startTime
	}
	endTime, err2 := time.ParseInLocation("2006-01-02 15:04:05", c.Query("endTime"), time.Local)
	if err2 == nil {
		param.EndTime = endTime
	}
	types := c.Query("types")
	param.Types = types

	if data, err := statisticsFacilityService.StatisticsFacilityByType(param); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(data, c)
	}

}
