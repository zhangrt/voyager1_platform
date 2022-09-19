package demo

import "github.com/zhangrt/voyager1_platform/repository"

type ServiceGroup struct {
	FacilityService
}

var (
	facilityRepository = repository.RepositoryGroupApp.DemoRepositoryGroup.FacilityRepository
)
