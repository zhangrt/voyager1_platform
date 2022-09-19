package repository

import (
	"github.com/zhangrt/voyager1_platform/repository/demo"
	"github.com/zhangrt/voyager1_platform/repository/statistics"
	"github.com/zhangrt/voyager1_platform/repository/test"
)

type RepositoryGroup struct {
	TestRepositoryGroup  test.RepositoryGroup
	DemoRepositoryGroup  demo.RepositoryGroup
	StatisticsRepository statistics.RepositoryGroup
}

var RepositoryGroupApp = new(RepositoryGroup)
