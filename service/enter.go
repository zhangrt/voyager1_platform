package service

import (
	"github.com/zhangrt/voyager1_platform/service/demo"
	"github.com/zhangrt/voyager1_platform/service/file"
	"github.com/zhangrt/voyager1_platform/service/statistics"
	"github.com/zhangrt/voyager1_platform/service/system"
	"github.com/zhangrt/voyager1_platform/service/test"
)

type ServiceGroup struct {
	TestServiceGroup   test.ServiceGroup
	FileServiceGroup   file.ServiceGroup
	SystemServiceGroup system.ServiceGroup
	DemoServiceGroup   demo.ServiceGroup
	StatisticsGroup    statistics.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
