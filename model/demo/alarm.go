package demo

import "github.com/zhangrt/voyager1_platform/global"

// 设备告警表 用于demo演示
type Alarm struct {
	global.GS_BASE_MODEL
}

func (Alarm) TableName() string {
	return "alarm"
}
