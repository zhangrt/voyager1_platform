package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1RoleMenu struct {
	global.GS_BASE_MODEL_ID_NONE
	MenuId string `json:"menuId" gorm:"column:vo1_menu_id;comment:菜单ID"`
	RoleId string `json:"-" gorm:"column:vo1_role_id;comment:角色ID"`
}

func (s Vo1RoleMenu) TableName() string {
	return "vo1_role_mtm_menu"
}
