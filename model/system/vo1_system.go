package system

import "github.com/zhangrt/voyager1_core/global"

// 配置文件结构体
type Vo1System struct {
	global.GS_BASE_MODEL_ID_STRING
	Name        string `json:"name" gorm:"comment:名称"`
	HomePage    string `json:"home_page" gorm:"comment:主页"`
	SerialNo    string `json:"serial_no" gorm:"comment:排序"`
	Description string `json:"description" gorm:"comment:描述"`
}

func (Vo1System) TableName() string {
	return "vo1_system"
}
