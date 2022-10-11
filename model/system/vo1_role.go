package system

import "github.com/zhangrt/voyager1_platform/global"

type Vo1Role struct {
	global.GS_BASE_MODEL
	RoleId        string        `json:"roleId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	RoleName      string        `json:"roleName" gorm:"comment:角色名"`                                    // 角色名
	ParentId      string        `json:"parentId" gorm:"comment:父角色ID"`                                  // 父角色ID
	DataRoleId    []*Vo1Role    `json:"dataRoleId,omitempty" gorm:"many2many:sys_data_authority_id;"`
	Children      []Vo1Role     `json:"children,omitempty" gorm:"-"`
	SysBaseMenus  []SysBaseMenu `json:"menus,omitempty" gorm:"many2many:vo1_role_mtm_menus;"`
	Persons       []Vo1Person   `json:"-" gorm:"many2many:vo1_person_mtm_role;"`
	DefaultRouter string        `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (Vo1Role) TableName() string {
	return "vo1_role"
}

func (Vo1Role) Init() Vo1Role {
	authority := Vo1Role{
		RoleId:        "",
		RoleName:      "",
		ParentId:      "",
		DataRoleId:    []*Vo1Role{},
		Children:      []Vo1Role{},
		SysBaseMenus:  []SysBaseMenu{},
		Persons:       []Vo1Person{},
		DefaultRouter: "",
	}

	return authority
}
