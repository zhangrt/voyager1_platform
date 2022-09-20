package system

import "time"

type SysUnit struct {
	CreatedAt  time.Time  // 创建时间
	UpdatedAt  time.Time  // 更新时间
	DeletedAt  *time.Time `sql:"index"`
	UnitId     string     `json:"unitId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 部门ID
	UnitName   string     `json:"unitName" gorm:"comment:角色名"`                                    // 部门名称
	ParentId   string     `json:"parentId" gorm:"comment:父角色ID"`                                  // 父级部门ID
	DataUnitId []*SysUnit `json:"dataUnitId,omitempty" gorm:"many2many:sys_data_unit_id;"`
	Children   []SysUnit  `json:"children,omitempty" gorm:"-"`
	Users      []SysUser  `json:"-" gorm:"many2many:sys_user_unit;"`
}

func (SysUnit) TableName() string {
	return "sys_unit"
}
