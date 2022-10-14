// 自动生成模板SysDictionary
package system

import (
	"github.com/zhangrt/voyager1_core/global"
)

// 如果含有time.Time 请自行import time包
type Vo1Dictionary struct {
	global.GS_BASE_MODEL_ID_STRING
	NameCN      string          `json:"nameCN" form:"nameCN" gorm:"column:name_cn;comment:字典名（中）"` // 字典名（中）
	NameEN      string          `json:"nameEN" form:"nameEN" gorm:"column:name_en;comment:字典名（英）"` // 字典名（英）
	ParentId    string          `json:"parentId" form:"parentId" gorm:"parent_id:name;"`
	SerialNo    string          `json:"SerialNo" form:"SerialNo" gorm:"column:serial_no;"`
	Description string          `json:"description" form:"description" gorm:"column:description;comment:描述"` // 描述
	Dictionarys []Vo1Dictionary `gorm:"-"`
}

func (Vo1Dictionary) TableName() string {
	return "vo1_dictionary"
}
