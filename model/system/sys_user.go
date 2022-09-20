package system

import (
	"github.com/zhangrt/voyager1_core/global"
)

type SysUser struct {
	global.GS_BASE_USER
	Authority   SysAuthority    `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority  `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Department  SysDepartment   `json:"department" gorm:"foreignKey:DepartmentId;references:DepartmentId;comment:用户部门"`
	Departments []SysDepartment `json:"departments"  gorm:"many2many:sys_user_department;"`
	Unit        SysUnit         `json:"unit" gorm:"foreignKey:UnitId;references:UnitId;comment:用户部门"`
	Units       []SysUnit       `json:"units"  gorm:"many2many:sys_user_unit;"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
