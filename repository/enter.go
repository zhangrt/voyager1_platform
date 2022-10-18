package repository

import (
	"github.com/zhangrt/voyager1_platform/repository/demo"
	"github.com/zhangrt/voyager1_platform/repository/statistics"
	"github.com/zhangrt/voyager1_platform/repository/test"
	"github.com/zhangrt/voyager1_platform/repository/voyager1"
)

type RepositoryGroup struct {
	TestRepositoryGroup  test.RepositoryGroup
	DemoRepositoryGroup  demo.RepositoryGroup
	StatisticsRepository statistics.RepositoryGroup
	Voyager1Repository   voyager1.RepositoryGroup
}

var RepositoryGroupApp = new(RepositoryGroup)
