package request

import (
	"time"

	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/demo"
)

type FacilitySearch struct {
	// 查询参数：设备名
	Name string `json:"name" form:"name"`
	// 查询参数：设备唯一标识码
	Code string `json:"code" form:"code"`
	// 查询参数：设备类型 用,隔开 0,1,2,3,4,5 <==> 电子围栏、火灾报警器、喇叭、CCTV、警报器、门禁
	Type string `json:"type" form:"type"`
	// 查询参数：设备状态 用,隔开 0,1,2 <==> 在线 离线 告警
	Status string `json:"status" form:"status"`
	// 查询参数：报警起始时间
	StartTime time.Time `json:"startTime" form:"startTime"`
	// 查询参数：报警结束
	EndTime time.Time `json:"endTime" form:"endTime"`
	request.PageInfo
}

type FacilityAdd struct {
	// 新增参数：设备名
	Name string `json:"name" form:"name"`
	// 新增参数：设备唯一标识码
	Code string `json:"code" form:"code"`
	// 新增参数：设备类型 0,1,2,3,4,5 <==> 电子围栏、火灾报警器、喇叭、CCTV、警报器、门禁
	Type string `json:"type" form:"type"`
	// 新增参数：设备状态 0,1,2 <==> 在线 离线 告警
	Status string `json:"status" form:"status"`
}

type FacilityUpdate struct {
	// 更新参数：主键ID
	ID uint `json:"id,string" form:"id"`
	// 更新参数：设备名
	Name string `json:"name" form:"name"`
	// 更新参数：设备唯一标识码
	Code string `json:"code" form:"code"`
	// 更新参数：设备类型 0,1,2,3,4,5 <==> 电子围栏、火灾报警器、喇叭、CCTV、警报器、门禁
	Type string `json:"type" form:"type"`
	// 更新参数：设备状态 0,1,2 <==> 在线 离线 告警
	Status string `json:"status" form:"status"`
	// 报警时间
	AlarmTime time.Time `json:"alarmTime" form:"alarmTime"`
}

func AddToFacility(add FacilityAdd) demo.Facility {
	var data demo.Facility
	data.Name = add.Name
	data.Code = add.Code
	data.Type = add.Type
	data.Status = add.Status
	return data
}

func UpdateToFacility(update FacilityUpdate) demo.Facility {
	var data demo.Facility
	data.ID = update.ID
	data.Name = update.Name
	data.Code = update.Code
	data.Type = update.Type
	data.Status = update.Status
	data.AlarmTime = update.AlarmTime
	return data
}
