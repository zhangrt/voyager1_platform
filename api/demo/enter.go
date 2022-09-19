package demo

import "github.com/zhangrt/voyager1_platform/service"

type ApiGroup struct {
	FacilityApi
}

var (
	facilityService = service.ServiceGroupApp.DemoServiceGroup.FacilityService
)
