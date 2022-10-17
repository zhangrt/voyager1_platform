package request

import (
	"github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_platform/model/system"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus  []system.Vo1Menu `json:"menus"`
	RoleId string           `json:"roleId"` // 角色ID
}

func DefaultMenu() []system.Vo1Menu {
	return []system.Vo1Menu{{
		GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "1"},
		ParentId:                "0",
		Url:                     "dashboard",
		Name:                    "仪表盘",
		Component:               "pages/dashboard/dashboard_view.dart",
		SerialNo:                1,
		Description:             "仪表盘",
		Icon:                    "dashboard",
	}}
}
