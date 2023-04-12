package voyager1

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system"
)

type Vo1PersonRepository struct {
}

func (t *Vo1PersonRepository) GetDepartmentsByPersonAccount(value string) (_ []system.Vo1Department, err error) {
	var department []system.Vo1Department
	var user system.Vo1Person
	// 这里需要保证不同用户之间account、phone、email都不相同，也不能存在A.account=B.phone的情况
	err = global.GS_DB.Where("account = ? or phone = ? or email = ?", value, value, value).First(&user).Error
	if err != nil {
		return *&department, err
	}
	err = global.GS_DB.Where("id in (?) and deleted <> 1", global.GS_DB.Where("vo1_person_id = ? and deleted <> 1", user.ID).Table("vo1_person_mtm_department").Select("vo1_department_id")).Find(&department).Error
	if err != nil {
		return *&department, err
	}
	return department, err

}
