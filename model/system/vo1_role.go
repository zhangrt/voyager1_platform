package system

import "github.com/zhangrt/voyager1_core/global"

type Vo1Role struct {
	global.GS_BASE_MODEL_ID_STRING
	Name        string      `json:"name" gorm:"comment:角色名"`       // 角色名
	OrganId     string      `json:"organId" gorm:"comment:组织机构ID"` // 所属组织机构ID，当organId值为空时代表该角色为公共角色
	SystemId    string      `json:"systemId" gorm:"系统ID"`
	Description string      `json:"description"`
	DataRoleId  []*Vo1Role  `json:"dataRoleId,omitempty" gorm:"many2many:sys_data_role_id;"`
	Children    []Vo1Role   `json:"children,omitempty" gorm:"-"`
	Vo1Menu     []Vo1Menu   `json:"menus,omitempty" gorm:"many2many:vo1_role_mtm_menus;"`
	Persons     []Vo1Person `json:"-" gorm:"many2many:vo1_person_mtm_role;"`
}

func (Vo1Role) TableName() string {
	return "vo1_role"
}

func (Vo1Role) Init() Vo1Role {
	authority := Vo1Role{
		GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{
			ID: "",
		},
		Name:       "",
		OrganId:    "",
		DataRoleId: []*Vo1Role{},
		Children:   []Vo1Role{},
		Vo1Menu:    []Vo1Menu{},
		Persons:    []Vo1Person{},
	}

	return authority
}
