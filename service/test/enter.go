package test

import "github.com/zhangrt/voyager1_platform/repository"

type ServiceGroup struct {
	TestService
}

var (
	testRepository = repository.RepositoryGroupApp.TestRepositoryGroup.TestRepository
)
