package request

import (
	model "github.com/zhangrt/voyager1_platform/model/system"
)

// User register structure
type Register struct {
	Account  string   `json:"account"` // 用户名，非必填
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Password string   `json:"passWord"`
	Name     string   `json:"name" gorm:"default:'Big Monster'"`
	Avatar   string   `json:"avatar" gorm:"default:'https://c-ssl.dtstatic.com/uploads/item/201901/19/20190119105005_uJPTs.thumb.1000_0.jpeg'"` // 头像
	RoleIds  []string `json:"roleIds"`
}

// User login structure
type Login struct {
	Identification string `json:"identification"` // 身份：用户名|手机号|邮箱
	Password       string `json:"password"`       // 密码
	OrganizationId string `json:"OrganizationId"`
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
	RoleId  string   `json:"roleId"`
	RoleIds []string `json:"roleIds"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID      string
	RoleIds []string `json:"roleIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID      string          `gorm:"primarykey"`                                                                                                                  // 主键ID
	Name    string          `json:"name" gorm:"default:系统用户;comment:用户昵称"`                                                                                       // 用户昵称
	Phone   string          `json:"phone"  gorm:"comment:用户手机号"`                                                                                                 // 用户角色ID
	RoleIds []string        `json:"roleIds" gorm:"-"`                                                                                                            // 角色ID
	Email   string          `json:"email"  gorm:"comment:用户邮箱"`                                                                                                  // 用户邮箱
	Avatar  string          `json:"avatar" gorm:"default:https://c-ssl.dtstatic.com/uploads/item/201901/19/20190119105005_uJPTs.thumb.1000_0.jpeg;comment:用户头像"` // 用户头像
	Roles   []model.Vo1Role `json:"-" gorm:"many2many:vo1_person_mtm_role;"`
}
