package system

type ServiceGroup struct {
	MenuService
	UserService
	InitDBService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	OperationRecordService
	WeatherService
}
