package voyager1

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type Vo1RoleRoter struct{}

func (s *Vo1RoleRoter) InitVo1RoleRouter(Router *gin.RouterGroup) {
	v1VoyagerRouter := Router.Group("v1/voyager1/role")
	roleApi := v1.ApiGroupApp.Voyager1Group.RoleApi
	{
		v1VoyagerRouter.POST("getMenusByRoleIds", roleApi.GetMenusByRoleIds)
	}

}
