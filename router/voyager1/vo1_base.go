package voyager1

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type Vo1BaseRouter struct{}

func (s *Vo1BaseRouter) InitVo1BaseRouter(Router *gin.RouterGroup) {
	v1VoyagerRouter := Router.Group("v1/voyager1")
	personApi := v1.ApiGroupApp.Voyager1Group.PersonApi
	{
		v1VoyagerRouter.POST("login", personApi.Login)
		v1VoyagerRouter.POST("register", personApi.Register)
		v1VoyagerRouter.POST("changePassword", personApi.ChangePassword)
		v1VoyagerRouter.POST("resetPassword", personApi.ResetPassword)
	}

}
