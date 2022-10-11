package system

import "time"

type SysOrganization struct {
	CreatedAt          time.Time          // 创建时间
	UpdatedAt          time.Time          // 更新时间
	DeletedAt          *time.Time         `sql:"index"`
	OrganizationId     string             `json:"organizationId" gorm:"not null;unique;primary_key;comment:组织机构ID;size:90"` // 组织机构ID
	OrganizationName   string             `json:"organizationName" gorm:"comment:组织机构名"`                                    // 组织机构名称
	ParentId           string             `json:"parentId" gorm:"comment:父组织机构ID"`                                          // 父级组织机构ID
	DataOrganizationId []*SysOrganization `json:"dataOrganizationId,omitempty" gorm:"many2many:sys_data_organization_id;"`
	Children           []SysOrganization  `json:"children,omitempty" gorm:"-"`
	Users              []Vo1Person        `json:"-" gorm:"many2many:sys_user_organization;"`
}

func (SysOrganization) TableName() string {
	return "vo1_organization"
}
