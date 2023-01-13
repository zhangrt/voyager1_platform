package voyager1

import "github.com/zhangrt/voyager1_platform/model/system"

type SystemApi struct{}

func (s *SystemApi) GetSystemModel(id string) (system system.Vo1System) {
	return
}

func (s *SystemApi) CreateSystemModel(systemReq system.Vo1System) (systemRes system.Vo1System) {
	return
}

func (s *SystemApi) UpdateSystemModel(id string, systemReq system.Vo1System) (systemRes system.Vo1System) {
	return
}

func (s *SystemApi) DeleteSystemModel(id string) (res bool) {
	return
}

func (s *SystemApi) GetSystemsRolesModel() (system *system.Vo1System) {
	return
}

func (s *SystemApi) GetSystemsRolesModelById(id string) (system *system.Vo1System) {
	return
}

func (s *SystemApi) UpdateSystemRolesModel(id string, roleSystemReq *system.Vo1RoleSystem) (roleSystemRes *system.Vo1RoleSystem) {
	return
}
