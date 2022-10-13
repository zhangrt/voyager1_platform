package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1PersonSystem struct {
	global.GS_BASE_MODEL_ID_STRING
	RoleId   uint `gorm:"column:role_id"`
	SystemId uint `gorm:"column:system_id"`
}

func (s *Vo1PersonSystem) TableName() string {
	return "vo1_person_mtm_system"
}
