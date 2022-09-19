package statistics

import "time"

type StatisticsFacilityByType struct {
	// 起始时间
	StartTime time.Time ` json:"startTime" `
	// 结束时间
	EndTime time.Time ` json:"endTime" `
	// 类型参数 0-5 用逗号隔开
	Types string ` json:"types" `
}

type StatisticsFacilityByStatus struct {
	StartTime time.Time
	EndTime   time.Time
	Facilitys []StatisticsFacilityStatus
}

type StatisticsFacilityType struct {
	Type string
	Num  int
}

type StatisticsFacilityStatus struct {
	Status string
	Num    int
}
