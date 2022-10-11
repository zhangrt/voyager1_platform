package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1Department struct {
	global.GS_BASE_MODEL
	Name             string
	OrganId          string
	ParentId         string `json:"parentId" gorm:"comment:父部门ID"` // 父级部门ID
	SerialNo         string
	Description      string
	DataDepartmentId []*Vo1Department `json:"dataDepartmentId,omitempty" gorm:"many2many:sys_data_department_id;"`
	Children         []Vo1Department  `json:"children,omitempty" gorm:"-"`
	Persons          []Vo1Person      `json:"-" gorm:"many2many:sys_user_department;"`
}

func (Vo1Department) TableName() string {
	return "vo1_department"
}
