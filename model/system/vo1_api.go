package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1Api struct {
	global.GS_BASE_MODEL
	Name        string
	Url         string
	Action      string
	ParentId    string
	SystemId    string
	SerialId    string
	Description string
}
