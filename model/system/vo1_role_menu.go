package system

type SysMenu struct {
	SysBaseMenu
	MenuId      string            `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string            `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu         `json:"children" gorm:"-"`
	Btns        map[string]string `json:"btns" gorm:"-"`
}

func (s SysMenu) TableName() string {
	return "vo1_role_mtm_menu"
}
