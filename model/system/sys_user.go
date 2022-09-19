package system

import (
	"github.com/zhangrt/voyager1_platform/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GS_BASE_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`                                                                                                     // 用户UUID
	Username    string         `json:"userName" gorm:"comment:用户登录名"`                                                                                                  // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                                                                       // 用户登录密码
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                                                      // 用户昵称
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户主题"`                                                                                      // 用户侧边主题
	HeaderImg   string         `json:"headerImg" gorm:"default:https://c-ssl.dtstatic.com/uploads/item/201901/19/20190119105005_uJPTs.thumb.1000_0.jpeg;comment:用户头像"` // 用户头像
	AuthorityId string         `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                                                                  // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string         `json:"phone"  gorm:"comment:用户手机号"` // 用户手机号
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`  // 用户邮箱
}

func (SysUser) TableName() string {
	return "sys_users"
}
