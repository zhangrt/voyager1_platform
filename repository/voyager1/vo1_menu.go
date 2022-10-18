package voyager1

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system"
	"gorm.io/gorm"
)

type Vo1RoleRepository struct {
}

func (t *Vo1RoleRepository) GetRolesMenusByRoleIds(ids []string) (_ []system.Vo1Role) {
	var roles []system.Vo1Role
	// 自定义sql预加载
	global.GS_DB.Preload("Vo1Menus", func(db *gorm.DB) *gorm.DB {
		return db.Order("serial_no")
	}).Where("id in ?", ids).Find(&roles)
	return roles
}
