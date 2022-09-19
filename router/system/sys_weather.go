package system

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type WeatherRouter struct{}

func (s *WeatherRouter) InitWeatherrRouter(Router *gin.RouterGroup) {
	weatherRouter := Router.Group("weather")
	weatherApi := v1.ApiGroupApp.SystemApiGroup.WeatherApi
	{
		weatherRouter.GET("getInfo", weatherApi.GetWeatherInfo)
	}
}
