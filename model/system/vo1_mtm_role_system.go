package system

type Vo1PersonSystem struct {
	RoleId   uint `gorm:"column:vo1_role_id"`
	SystemId uint `gorm:"column:vo1_system_id"`
}

func (s *Vo1PersonSystem) TableName() string {
	return "vo1_person_mtm_system"
}
