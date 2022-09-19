package demo

import (
	"time"

	"github.com/zhangrt/voyager1_platform/global"
)

// 基础设备表 用于demo演示
type Facility struct {
	global.GS_BASE_MODEL
	// 设备名
	Name string `json:"name" form:"name" gorm:"column:name;size:255;comment:设备名"`
	// 设备唯一标识码
	Code string `json:"code" form:"code" gorm:"column:code;size:128;index:idx_facility_code,unique;comment:设备唯一标识码"`
	// 设备类型 0 1 2 3 4 5 <==> 电子围栏、火灾报警器、喇叭、CCTV、警报器、门禁
	Type string `json:"type" form:"type" gorm:"column:type;size:1;comment:设备类型 0 1 2 3 4 5 <==> 电子围栏、火灾报警器、喇叭、CCTV、警报器、门禁"`
	// 设备状态 0 1 2 <==> 在线 离线 告警
	Status string `json:"status" form:"status" gorm:"column:status;size:1;comment:设备状态0 1 2 <==> 在线 离线 告警"`
	// 报警时间
	AlarmTime time.Time `json:"alarmTime" form:"alarmTime" gorm:"column:alarm_time; comment:报警时间"`
}

func (Facility) TableName() string {
	return "facility"
}
