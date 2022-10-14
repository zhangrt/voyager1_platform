package response

import "github.com/zhangrt/voyager1_platform/model/system"

type Vo1MenusResponse struct {
	Menus []system.Vo1Menu `json:"menus"`
}

type Vo1MenuResponse struct {
	Menu system.Vo1Menu `json:"menu"`
}
