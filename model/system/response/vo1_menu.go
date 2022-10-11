package response

import "github.com/zhangrt/voyager1_platform/model/system"

type SysMenusResponse struct {
	Menus []system.Vo1Menu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []system.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu system.SysBaseMenu `json:"menu"`
}
