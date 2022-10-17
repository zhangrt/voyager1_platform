package service

import (
	"github.com/zhangrt/voyager1_platform/service/demo"
	"github.com/zhangrt/voyager1_platform/service/file"
	"github.com/zhangrt/voyager1_platform/service/statistics"
	"github.com/zhangrt/voyager1_platform/service/system"
	"github.com/zhangrt/voyager1_platform/service/test"
	"github.com/zhangrt/voyager1_platform/service/voyager1"
)

type ServiceGroup struct {
	TestServiceGroup   test.ServiceGroup
	FileServiceGroup   file.ServiceGroup
	SystemServiceGroup system.ServiceGroup
	DemoServiceGroup   demo.ServiceGroup
	StatisticsGroup    statistics.ServiceGroup
	Voyager1Group      voyager1.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
