package system

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/response"
	systemRes "github.com/zhangrt/voyager1_platform/model/system/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WeatherApi struct{}

// @Tags Weather
// @Summary 天气信息
// @Produce  application/json
// @Param cityCode query string true "城市编码cityCode 101010100 "
// @Success 200 {object} response.Response{data=systemRes.Vo1WeatherInfo,msg=string} "通过城市编码获取天气信息"
// @Router /weather/getInfo [get]
func (w *WeatherApi) GetWeatherInfo(c *gin.Context) {

	cityCode := c.Query("cityCode")

	if get_back, err := weatherService.GetWeatherInfo(cityCode); err != nil {
		global.GS_LOG.Error("get failed", zap.Error(err))
		response.OkWithDetailed(systemRes.Vo1WeatherInfo{
			Code:        "101010100",
			Temperature: "33.5",
			Location:    "北京",
			Humidness:   "51%",
			Info:        "晴转多云",
			Time:        "17:55",
		}, "测试数据", c)
	} else {
		response.OkWithDetailed(get_back, "获取成功", c)
	}

}
