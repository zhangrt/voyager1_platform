package system

type RouterGroup struct {
	BaseRouter
	JwtRouter
	SysRouter
	InitRouter
	MenuRouter
	UserRouter
	CasbinRouter
	AuthorityRouter
	OperationRecordRouter
	DictionaryRouter
	DictionaryDetailRouter
	AuthorityBtnRouter
	WeatherRouter
}
