package system

import (
	"errors"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system"

	"gorm.io/gorm"
)

type BaseMenuService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error

func (baseMenuService *BaseMenuService) DeleteBaseMenu(id string) (err error) {
	err = global.GS_DB.Preload("MenuBtn").Preload("Parameters").Where("parent_id = ?", id).First(&system.Vo1Menu{}).Error
	if err != nil {
		var menu system.Vo1Menu
		db := global.GS_DB.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		if err != nil {
			global.GS_LOG.Error(err.Error())
		}
		if len(menu.Roles) > 0 {
			err = global.GS_DB.Model(&menu).Association("SysAuthoritys").Delete(&menu.Roles)
		} else {
			err = db.Error
			if err != nil {
				return
			}
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBaseMenu
//@description: 更新路由
//@param: menu model.Vo1Menu
//@return: err error

func (baseMenuService *BaseMenuService) UpdateBaseMenu(menu system.Vo1Menu) (err error) {
	var oldMenu system.Vo1Menu
	upDateMap := make(map[string]interface{})
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["url"] = menu.Url
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["description"] = menu.Description
	upDateMap["icon"] = menu.Icon
	upDateMap["serial_no"] = menu.SerialNo

	err = global.GS_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&system.Vo1Menu{}).Error, gorm.ErrRecordNotFound) {
				global.GS_LOG.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}

		txErr := db.Updates(upDateMap).Error
		if txErr != nil {
			global.GS_LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id float64
//@return: menu system.Vo1Menu, err error

func (baseMenuService *BaseMenuService) GetBaseMenuById(id string) (menu system.Vo1Menu, err error) {
	err = global.GS_DB.Preload("MenuBtn").Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
