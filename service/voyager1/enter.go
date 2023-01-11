package voyager1

import "github.com/zhangrt/voyager1_platform/repository"

type ServiceGroup struct {
	PersonService
	RoleService
	MenuService
}

var (
	roleRepository   = repository.RepositoryGroupApp.Voyager1Repository.Vo1RoleRepository
	personRepository = repository.RepositoryGroupApp.Voyager1Repository.Vo1PersonRepository
)
