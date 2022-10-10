package system

import "github.com/zhangrt/voyager1_platform/global"

type SysAuthority struct {
	global.GS_BASE_MODEL
	RoleId        string          `json:"roleId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	RoleName      string          `json:"roleName" gorm:"comment:角色名"`                                    // 角色名
	ParentId      string          `json:"parentId" gorm:"comment:父角色ID"`                                  // 父角色ID
	DataRoleId    []*SysAuthority `json:"dataRoleId,omitempty" gorm:"many2many:sys_data_authority_id;"`
	Children      []SysAuthority  `json:"children,omitempty" gorm:"-"`
	SysBaseMenus  []SysBaseMenu   `json:"menus,omitempty" gorm:"many2many:vo1_role_mtm_menus;"`
	Users         []SysUser       `json:"-" gorm:"many2many:vo1_person_mtm_role;"`
	DefaultRouter string          `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (SysAuthority) TableName() string {
	return "vo1_role"
}

func (SysAuthority) Init() SysAuthority {
	authority := SysAuthority{
		RoleId:        "",
		RoleName:      "",
		ParentId:      "",
		DataRoleId:    []*SysAuthority{},
		Children:      []SysAuthority{},
		SysBaseMenus:  []SysBaseMenu{},
		Users:         []SysUser{},
		DefaultRouter: "",
	}

	return authority
}
