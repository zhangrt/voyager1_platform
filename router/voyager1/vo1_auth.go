package voyager1

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type Vo1AuthRouter struct{}

func (s *Vo1AuthRouter) InitVo1AuthRouter(Router *gin.RouterGroup) {
	v1VoyagerRouter := Router.Group("v1/voyager1/auth")
	casbinApi := v1.ApiGroupApp.Voyager1Group.CasbinApi
	jwtApi := v1.ApiGroupApp.Voyager1Group.JwtApi
	roleApi := v1.ApiGroupApp.Voyager1Group.RoleApi
	{
		v1VoyagerRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
		v1VoyagerRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist)
		v1VoyagerRouter.POST("getMenusByRoleIds", roleApi.GetMenusByRoleIds)
	}

}
