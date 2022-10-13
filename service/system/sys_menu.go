package system

import (
	"errors"
	"strconv"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/system"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: treeMap map[string][]system.Vo1Menu, err error

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) getMenuTreeMap(authorityIds []string) (treeMap map[string][]system.Vo1Menu, err error) {
	var allMenus []system.Vo1Menu
	treeMap = make(map[string][]system.Vo1Menu)
	err = global.GS_DB.Where("authority_id in ?", authorityIds).Order("sort").Preload("Parameters").Find(&allMenus).Error
	if err != nil {
		return
	}
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: menus []system.Vo1Menu, err error

func (menuService *MenuService) GetMenuTree(authorityIds []string) (menus []system.Vo1Menu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityIds)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *model.Vo1Menu, treeMap map[string][]model.Vo1Menu
//@return: err error

func (menuService *MenuService) getChildrenList(menu *system.Vo1Menu, treeMap map[string][]system.Vo1Menu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetInfoList
//@description: 获取路由分页
//@return: list interface{}, total int64,err error

func (menuService *MenuService) GetInfoList() (list interface{}, total int64, err error) {
	var menuList []system.Vo1Menu
	treeMap, err := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.Vo1Menu, treeMap map[string][]model.Vo1Menu
//@return: err error

func (menuService *MenuService) getBaseChildrenList(menu *system.Vo1Menu, treeMap map[string][]system.Vo1Menu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.SerialNo))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu model.Vo1Menu
//@return: error

func (menuService *MenuService) AddBaseMenu(menu system.Vo1Menu) error {
	if !errors.Is(global.GS_DB.Where("name = ?", menu.Name).First(&system.Vo1Menu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.GS_DB.Create(&menu).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: treeMap map[string][]system.Vo1Menu, err error

func (menuService *MenuService) getBaseMenuTreeMap() (treeMap map[string][]system.Vo1Menu, err error) {
	var allMenus []system.Vo1Menu
	treeMap = make(map[string][]system.Vo1Menu)
	err = global.GS_DB.Order("sort").Preload("MenuBtn").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: menus []system.Vo1Menu, err error

func (menuService *MenuService) GetBaseMenuTree() (menus []system.Vo1Menu, err error) {
	treeMap, err := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.Vo1Menu, authorityId string
//@return: err error

func (menuService *MenuService) AddMenuAuthority(menus []system.Vo1Menu, authorityId string) (err error) {
	var auth system.Vo1Role
	auth.ID = authorityId
	auth.Vo1Menu = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: menus []system.Vo1Menu, err error

func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (menus []system.Vo1Menu, err error) {
	err = global.GS_DB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	// sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	// err = global.GS_DB.Raw(sql, authorityId).Scan(&menus).Error
	return menus, err
}
