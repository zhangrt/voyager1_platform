package test

import (
	"github.com/zhangrt/voyager1_platform/global"
)

type Test struct {
	global.GS_BASE_MODEL
	TestName string `json:"testName" form:"testName" gorm:"column:test_name;size:255;index:idx_name,unique;comment:测试名字，唯一"`
}

// 可指定具体表名 解决表名复数问题
func (Test) TableName() string {
	return "test"
}
