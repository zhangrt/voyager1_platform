package system

import (
	"github.com/zhangrt/voyager1_core/global"
)

// type Vo1Menu struct {
// 	SysBaseMenu
// 	// MenuId   string    `json:"menuId" gorm:"comment:菜单ID"`
// 	OrganId  string    `json:"organId" gorm:"comment:角色ID"`
// 	RoleId   string    `json:"-" gorm:"comment:角色ID"`
// 	Children []Vo1Menu `json:"children" gorm:"-"`
// }

type Vo1Menu struct {
	global.GS_BASE_MODEL_ID_STRING
	// MenuId        uint                   `gorm:"column:menu_id" json:"menu_id,string" form:"menu_id"` //菜单ID
	// OrganId       string    `json:"organId" gorm:"comment:组织机构ID"`
	RoleId        string    `json:"-" gorm:"comment:角色ID"`
	MenuLevel     uint      `json:"-"`
	ParentId      string    `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Url           string    `json:"url" gorm:"comment:路由path"`         // 路由path
	Name          string    `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden        bool      `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component     string    `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	SerialNo      int       `json:"serialNo" gorm:"comment:排序标记"`      // 排序标记 0、1、2、3、4、5 ...
	Description   string    `json:"description" gorm:"comment:描述"`     // 描述
	Icon          string    `json:"icon" gorm:"comment:图标"`            // 图标
	SysAuthoritys []Vo1Role `json:"authoritys" gorm:"many2many:vo1_role_mtm_menu;"`
	Children      []Vo1Menu `json:"children" gorm:"-"`
}

func (Vo1Menu) TableName() string {
	return "vo1_menu"
}
