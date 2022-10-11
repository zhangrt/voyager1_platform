// 自动生成模板SysDictionary
package system

import (
	"github.com/zhangrt/voyager1_core/global"
)

// 如果含有time.Time 请自行import time包
type Vo1Dictionary struct {
	global.GS_BASE_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"` // 字典名（中）
	ParentId    string
	SerialNo    string
	Type        string `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`                  // 字典名（英）
	Status      *bool  `json:"status" form:"status" gorm:"column:status;comment:状态"`                // 状态
	Value       string `json:"value" form:"value" gorm:"column:value;comment:字典值"`                  // 字典值
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述"` // 描述
	Dictionarys []Vo1Dictionary
}

func (Vo1Dictionary) TableName() string {
	return "vo1_dictionary"
}
