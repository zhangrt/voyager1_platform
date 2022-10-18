package v1

import (
	"github.com/zhangrt/voyager1_platform/api/demo"
	"github.com/zhangrt/voyager1_platform/api/file"
	"github.com/zhangrt/voyager1_platform/api/statistics"
	"github.com/zhangrt/voyager1_platform/api/system"
	test "github.com/zhangrt/voyager1_platform/api/test"
	voyager1 "github.com/zhangrt/voyager1_platform/api/v1/voyager1"
)

// API分组
type ApiGroup struct {
	TestApiGroup       test.ApiGroup
	FileApiGroup       file.ApiGroup
	SystemApiGroup     system.ApiGroup
	DemoApiGroup       demo.ApiGroup
	StatisticsApiGroup statistics.ApiGroup
	Voyager1Group      voyager1.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
