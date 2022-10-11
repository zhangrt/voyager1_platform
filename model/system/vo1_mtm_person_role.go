package system

type Vo1PersonRole struct {
	PersonId uint   `gorm:"column:vo1_user_id"`
	RoleId   string `gorm:"column:vo1_role_id"`
}

func (s *Vo1PersonRole) TableName() string {
	return "vo1_person_mtm_role"
}
