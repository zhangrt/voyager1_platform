package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1PersonRole struct {
	global.GS_BASE_MODEL_ID_NONE
	PersonId string `gorm:"column:vo1_person_id"`
	RoleId   string `gorm:"column:vo1_role_id"`
}

func (s *Vo1PersonRole) TableName() string {
	return "vo1_person_mtm_role"
}
