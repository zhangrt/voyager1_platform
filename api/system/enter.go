package system

import "github.com/zhangrt/voyager1_platform/service"

type ApiGroup struct {
	DBApi
	JwtApi
	SystemApi
	CasbinApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
	OperationRecordApi
	DictionaryDetailApi
	UserApi
	WeatherApi
}

var (
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	weatherService          = service.ServiceGroupApp.SystemServiceGroup.WeatherService
)
