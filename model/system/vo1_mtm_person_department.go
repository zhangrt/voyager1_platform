package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1PersonDepartment struct {
	global.GS_BASE_MODEL_ID_NONE
	PersonId     uint   `gorm:"column:vo1_person_id"`
	DepartmentId string `gorm:"column:vo1_department_id"`
}

func (s *Vo1PersonDepartment) TableName() string {
	return "vo1_person_mtm_department"
}
