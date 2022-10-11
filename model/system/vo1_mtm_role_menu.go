package system

type Vo1RoleMenu struct {
	MenuId      string `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string `json:"-" gorm:"comment:角色ID"`
}

func (s Vo1RoleMenu) TableName() string {
	return "vo1_role_mtm_menu"
}
