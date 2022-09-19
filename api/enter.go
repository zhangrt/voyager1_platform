package v1

import (
	"github.com/zhangrt/voyager1_platform/api/demo"
	"github.com/zhangrt/voyager1_platform/api/file"
	"github.com/zhangrt/voyager1_platform/api/statistics"
	"github.com/zhangrt/voyager1_platform/api/system"
	test "github.com/zhangrt/voyager1_platform/api/test"
)

// API分组
type ApiGroup struct {
	TestApiGroup       test.ApiGroup
	FileApiGroup       file.ApiGroup
	SystemApiGroup     system.ApiGroup
	DemoApiGroup       demo.ApiGroup
	StatisticsApiGroup statistics.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
