package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1Api struct {
	global.GS_BASE_MODEL_ID_STRING
	Name        string
	Url         string
	Action      string
	ParentId    string
	SystemId    string
	SerialId    string
	Description string
}

func (Vo1Api) TableName() string {
	return "vo1_api"
}
