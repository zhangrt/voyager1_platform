package voyager1

import "github.com/zhangrt/voyager1_platform/service"

type ApiGroup struct {
	PersonApi
	RoleApi
	JwtApi
	CasbinApi
	SystemApi
}

var (
	personService = service.ServiceGroupApp.Voyager1Group.PersonService
	roleService   = service.ServiceGroupApp.Voyager1Group.RoleService
	menuService   = service.ServiceGroupApp.Voyager1Group.MenuService
	systemService = service.ServiceGroupApp.Voyager1Group.SystemService
)
