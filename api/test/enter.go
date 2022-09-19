package test

import "github.com/zhangrt/voyager1_platform/service"

type ApiGroup struct {
	TestApi
}

var (
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService

	testService = service.ServiceGroupApp.TestServiceGroup.TestService
)
