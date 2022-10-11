package request

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus"`
	AuthorityId string               `json:"authorityId"` // 角色ID
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		GS_BASE_MODEL: global.GS_BASE_MODEL{ID: 1},
		ParentId:      "0",
		Path:          "dashboard",
		Name:          "dashboard",
		Component:     "pages/dashboard/dashboard_view.dart",
		Sort:          1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
