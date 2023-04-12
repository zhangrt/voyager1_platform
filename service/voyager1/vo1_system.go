package voyager1

import "github.com/zhangrt/voyager1_platform/model/system"

type SystemService struct{}

func (s *SystemService) GetSystemModel(id string) (system system.Vo1System) {
	return
}

func (s *SystemService) CreateSystemModel(systemReq system.Vo1System) (systemRes system.Vo1System) {
	return
}

func (s *SystemService) UpdateSystemModel(id string, systemReq system.Vo1System) (systemRes system.Vo1System) {
	return
}

func (s *SystemService) DeleteSystemModel(id string) (res bool) {
	return
}

func (s *SystemService) GetSystemsRolesModel() (system *system.Vo1System) {
	return
}

func (s *SystemService) GetSystemsRolesModelById(id string) (system *system.Vo1System) {
	return
}

func (s *SystemService) UpdateSystemRolesModel(id string, roleSystemReq *system.Vo1RoleSystem) (roleSystemRes *system.Vo1RoleSystem) {
	return
}
