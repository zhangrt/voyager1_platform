package voyager1

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//@function: Register
//@description: 用户注册
//@param: u model.Vo1Person
//@return: userInter system.Vo1Person, err error

type PersonService struct{}

func (ps *PersonService) Register(u system.Vo1Person) (userInter system.Vo1Person, err error) {
	var user system.Vo1Person
	if !errors.Is(global.GS_DB.Where("account = ?", u.Account).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.ID = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	err = global.GS_DB.Create(&u).Error
	return u, err
}

//@function: Login
//@description: 用户登录
//@param: u *model.Vo1Person
//@return: err error, userInter *model.Vo1Person

func (ps *PersonService) Login(u *system.Vo1Person) (userInter *system.Vo1Person, err error) {
	if nil == global.GS_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.Vo1Person
	// 这里需要保证不同用户之间account、phone、email都不相同，也不能存在A.account=B.phone的情况
	err = global.GS_DB.Where("account = ? or phone = ? or email = ?", u.Account, u.Phone, u.Email).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		// var am system.Vo1Menu
		// ferr := global.GS_DB.First(&am, "name = ? AND role_id = ?", user.Role.DefaultRouter, user.RoleId).Error
		// if errors.Is(ferr, gorm.ErrRecordNotFound) {
		// 	user.Role.DefaultRouter = "404"
		// }

		// 登录成功这里查询该用户跟组织机构相关的角色信息
		var roles []system.Vo1Role
		if u.OrganizationId == "" {
			// select * from role where id in (select role_id from vo1_person_mtm_role where person_id = ?)
			global.GS_DB.Where("id in (?)", global.GS_DB.Table("vo1_person_mtm_role").Select("vo1_role_id").Where("vo1_person_id = ?", user.ID)).Find(&roles)
		} else {
			// select * from role where organ_id = ?
			global.GS_DB.Where("organ_id = ?", u.OrganizationId).Find(&roles)
		}

		user.Roles = roles

		var ids []string
		for _, r := range roles {
			ids = append(ids, r.ID)
		}

		user.RoleIds = ids

		var updateUser system.Vo1Person
		updateUser.ID = user.ID
		// 设置登录时间
		global.GS_DB.Model(&updateUser).Update("last_login_time", time.Now())

	}

	return &user, err
}

//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.Vo1Person, newPassword string
//@return: userInter *model.Vo1Person,err error

func (ps *PersonService) ChangePassword(u *system.Vo1Person, newPassword string) (userInter *system.Vo1Person, err error) {
	var user system.Vo1Person
	err = global.GS_DB.Where("account = ?", u.Account).First(&user).Error
	if err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GS_DB.Save(&user).Error
	return &user, err

}

//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (ps *PersonService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GS_DB.Model(&system.Vo1Person{})
	var userList []system.Vo1Person
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (ps *PersonService) SetUserAuthorities(id string, authorityIds []string) (err error) {
	return global.GS_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.Vo1PersonRole{}, "vo1_person_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []system.Vo1PersonRole{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.Vo1PersonRole{
				PersonId: id, RoleId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("vo1_user_id = ?", id).First(&system.Vo1Person{}).Update("role_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (ps *PersonService) DeleteUser(id string) (err error) {
	var user system.Vo1Person
	err = global.GS_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.GS_DB.Delete(&[]system.Vo1PersonRole{}, "vo1_person_id = ?", id).Error
	return err
}

//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.Vo1Person
//@return: err error, user model.Vo1Person

func (ps *PersonService) SetUserInfo(req system.Vo1Person) error {
	return global.GS_DB.Updates(&req).Error
}

//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.Vo1Person

func (ps *PersonService) GetUserInfo(uuid string) (user system.Vo1Person, err error) {
	var reqUser system.Vo1Person
	err = global.GS_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}

	return reqUser, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.Vo1Person

func (ps *PersonService) FindUserById(id int) (user *system.Vo1Person, err error) {
	var u system.Vo1Person
	err = global.GS_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.Vo1Person

func (ps *PersonService) FindUserByUuid(uuid string) (user *system.Vo1Person, err error) {
	var u system.Vo1Person
	if err = global.GS_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (ps *PersonService) ResetPassword(ID string) (err error) {
	err = global.GS_DB.Model(&system.Vo1Person{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

func (ps *PersonService) GetDepartmentsByAccountOrPhoneOrEmail(value string) (departments []system.Vo1Department, err error) {

	department, err := personRepository.GetDepartmentsByPersonAccount(value)

	return *&department, err
}
