package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1PersonSystem struct {
	global.GS_BASE_MODEL_ID_NONE
	RoleId   uint `gorm:"column:vo1_role_id"`
	SystemId uint `gorm:"column:vo1_system_id"`
}

func (s *Vo1PersonSystem) TableName() string {
	return "vo1_person_mtm_system"
}
