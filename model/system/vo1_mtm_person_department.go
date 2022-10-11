package system

type Vo1PersonDepartment struct {
	PersonId     uint   `gorm:"column:vo1_person_id"`
	DepartmentId string `gorm:"column:vo1_department_id"`
}

func (s *Vo1PersonDepartment) TableName() string {
	return "vo1_person_mtm_department"
}
