package system

import (
	"errors"
	"strconv"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/model/system/response"

	authentation "github.com/zhangrt/voyager1_core/auth/luna"

	"gorm.io/gorm"
)

var ErrRoleExistence = errors.New("存在相同角色id")

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.Vo1Role
//@return: authority system.Vo1Role, err error

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth system.Vo1Role) (authority system.Vo1Role, err error) {
	var authorityBox system.Vo1Role
	if !errors.Is(global.GS_DB.Where("id = ?", auth.ID).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	err = global.GS_DB.Create(&auth).Error
	return auth, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: authority system.Vo1Role, err error

func (authorityService *AuthorityService) CopyAuthority(copyInfo response.Vo1RoleCopyResponse) (authority system.Vo1Role, err error) {
	var authorityBox system.Vo1Role
	if !errors.Is(global.GS_DB.Where("id = ?", copyInfo.Role.ID).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}
	copyInfo.Role.Children = []system.Vo1Role{}
	menus, err := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldRoleId})
	if err != nil {
		return
	}
	var baseMenu []system.Vo1Menu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.ID)
		v.ID = string(intNum)
		baseMenu = append(baseMenu, v)
	}
	copyInfo.Role.Vo1Menu = baseMenu
	err = global.GS_DB.Create(&copyInfo.Role).Error
	if err != nil {
		return
	}

	auth := authentation.NewCasbin()
	paths := auth.GetPolicyPathByAuthorityId(copyInfo.OldRoleId)
	err = auth.UpdateCasbin(copyInfo.Role.ID, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Role)
	}
	return copyInfo.Role, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.Vo1Role
//@return: authority system.Vo1Role, err error

func (authorityService *AuthorityService) UpdateAuthority(auth system.Vo1Role) (authority system.Vo1Role, err error) {
	err = global.GS_DB.Where("id = ?", auth.ID).First(&system.Vo1Role{}).Updates(&auth).Error
	return auth, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRole
//@description: 删除角色
//@param: auth *model.Vo1Role
//@return: err error

func (authorityService *AuthorityService) DeleteAuthority(auth *system.Vo1Role) (err error) {
	if errors.Is(global.GS_DB.Debug().Preload("Persons").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Persons) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GS_DB.Where("id = ?", auth.ID).First(&system.Vo1Person{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GS_DB.Where("parent_id = ?", auth.ID).First(&system.Vo1Role{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.GS_DB.Preload("Vo1Menus").Where("id = ?", auth.ID).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.Vo1Menu) > 0 {
		err = global.GS_DB.Model(auth).Association("Vo1Menus").Delete(auth.Vo1Menu)
		if err != nil {
			return
		}
		// err = db.Association("Vo1Menus").Delete(&auth)
	} else {
		err = db.Error
		if err != nil {
			return
		}
	}
	err = global.GS_DB.Delete(&[]system.Vo1PersonRole{}, "vo1_role_id = ?", auth.ID).Error
	if err != nil {
		global.GS_LOG.Error(err.Error())
	}

	casbin := authentation.NewCasbin()
	casbin.ClearCasbin(0, auth.ID)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GS_DB.Model(&system.Vo1Role{})
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	var authority []system.Vo1Role
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return authority, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.Vo1Role
//@return: sa system.Vo1Role, err error

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.Vo1Role) (sa system.Vo1Role, err error) {
	err = global.GS_DB.Preload("DataAuthorityId").Where("id = ?", auth.ID).First(&sa).Error
	return sa, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.Vo1Role
//@return: error

func (authorityService *AuthorityService) SetDataAuthority(auth system.Vo1Role) error {
	var s system.Vo1Role
	global.GS_DB.Preload("DataAuthorityId").First(&s, "id = ?", auth.ID)
	err := global.GS_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataRoleId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.Vo1Role
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *system.Vo1Role) error {
	var s system.Vo1Role
	global.GS_DB.Preload("Vo1Menus").First(&s, "id = ?", auth.ID)
	err := global.GS_DB.Model(&s).Association("Vo1Menus").Replace(&auth.Vo1Menu)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.Vo1Role
//@return: err error

func (authorityService *AuthorityService) findChildrenAuthority(authority *system.Vo1Role) (err error) {
	err = global.GS_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.ID).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
