package system

import (
	"github.com/zhangrt/voyager1_core/global"
)

type Vo1Person struct {
	global.GS_BASE_USER
	Role  Vo1Role   `json:"role" gorm:"foreignKey:RoleId;references:RoleId;comment:用户角色"`
	Roles []Vo1Role `json:"roles" gorm:"many2many:vo1_person_mtm_role;"`
	// Department  SysDepartment   `json:"department" gorm:"foreignKey:DepartmentId;references:DepartmentId;comment:用户部门"`
	// Departments []SysDepartment `json:"departments"  gorm:"many2many:sys_user_department;"`
	// Unit        SysUnit         `json:"unit" gorm:"foreignKey:UnitId;references:UnitId;comment:用户部门"`
	// Units       []SysUnit       `json:"units"  gorm:"many2many:sys_user_unit;"`
}

func (Vo1Person) TableName() string {
	return "vo1_person"
}
