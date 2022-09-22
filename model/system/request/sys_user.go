package request

import model "github.com/zhangrt/voyager1_platform/model/system"

// User register structure
type Register struct {
	Account      string   `json:"account"`
	Password     string   `json:"passWord"`
	Name         string   `json:"name" gorm:"default:'Big Monster'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'https://c-ssl.dtstatic.com/uploads/item/201901/19/20190119105005_uJPTs.thumb.1000_0.jpeg'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:888"`
	AuthorityIds []string `json:"authorityIds"`
}

// User login structure
type Login struct {
	Account  string `json:"account"`  // 用户名
	Password string `json:"password"` // 密码
	// Captcha   string `json:"captcha"`   // 验证码
	// CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordStruct struct {
	Account     string `json:"account"`     // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                 `gorm:"primarykey"`                                                                                                                     // 主键ID
	Name         string               `json:"name" gorm:"default:系统用户;comment:用户昵称"`                                                                                          // 用户昵称
	Phone        string               `json:"phone"  gorm:"comment:用户手机号"`                                                                                                    // 用户角色ID
	AuthorityIds []string             `json:"authorityIds" gorm:"-"`                                                                                                          // 角色ID
	Email        string               `json:"email"  gorm:"comment:用户邮箱"`                                                                                                     // 用户邮箱
	HeaderImg    string               `json:"headerImg" gorm:"default:https://c-ssl.dtstatic.com/uploads/item/201901/19/20190119105005_uJPTs.thumb.1000_0.jpeg;comment:用户头像"` // 用户头像
	SideMode     string               `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                                                                // 用户侧边主题
	Authorities  []model.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
