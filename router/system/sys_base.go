package system

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		baseRouter.POST("login", userApi.Login)
	}

}
