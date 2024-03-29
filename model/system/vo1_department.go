package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1Department struct {
	global.GS_BASE_MODEL_ID_STRING
	Name        string          `json:"name" gorm:"comment:name"`
	OrganId     string          `json:"organId" gorm:"comment:organId"`
	ParentId    string          `json:"parentId" gorm:"comment:父部门ID"` // 父级部门ID
	SerialNo    string          `json:"serialNo" gorm:"comment:serialNo"`
	Description string          `json:"description" gorm:"comment:description"`
	Children    []Vo1Department `json:"children,omitempty" gorm:"-"`
	Persons     []Vo1Person     `json:"-" gorm:"many2many:vo1_person_mtm_department;"`
}

func (Vo1Department) TableName() string {
	return "vo1_department"
}
