package voyager1

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system"
	"gorm.io/gorm"
)

type Vo1PersonRepository struct {
}

func (t *Vo1PersonRepository) GetDepartmentsByPersonAccount(value string) (_ system.Vo1Person) {
	var persons system.Vo1Person
	// 自定义sql预加载
	global.GS_DB.Preload("Departments", func(db *gorm.DB) *gorm.DB {
		return db.Order("serial_no")
	}).Where("account = ? or phone = ? or email = ?", value, value, value).First(&persons)
	return persons
}
