package system

import "time"

type SysDepartment struct {
	CreatedAt        time.Time        // 创建时间
	UpdatedAt        time.Time        // 更新时间
	DeletedAt        *time.Time       `sql:"index"`
	DepartmentId     string           `json:"departmentId" gorm:"not null;unique;primary_key;comment:部门ID;size:90"` // 部门ID
	DepartmentName   string           `json:"departmentName" gorm:"comment:部门名"`                                    // 部门名称
	ParentId         string           `json:"parentId" gorm:"comment:父部门ID"`                                        // 父级部门ID
	DataDepartmentId []*SysDepartment `json:"dataDepartmentId,omitempty" gorm:"many2many:sys_data_department_id;"`
	Children         []SysDepartment  `json:"children,omitempty" gorm:"-"`
	Users            []SysUser        `json:"-" gorm:"many2many:sys_user_department;"`
}

func (SysDepartment) TableName() string {
	return "vo1_department"
}
