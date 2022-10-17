package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1Organization struct {
	global.GS_BASE_MODEL_ID_STRING
	Name        string            `json:"name" gorm:"comment:组织机构名"`       // 组织机构名称
	ParentId    string            `json:"parentId" gorm:"comment:父组织机构ID"` // 父级组织机构ID
	SerialNo    int               `json:"serialNo"`
	Description string            `json:"description"`
	Children    []Vo1Organization `json:"children,omitempty" gorm:"-"`
	Users       []Vo1Person       `json:"-" gorm:"many2many:sys_user_organization;"`
}

func (Vo1Organization) TableName() string {
	return "vo1_organization"
}
