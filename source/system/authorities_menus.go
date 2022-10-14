package system

import (
	"context"

	sysModel "github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/service/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenuAuthority = initOrderMenu + initOrderAuthority

type initMenuAuthority struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenuAuthority, &initMenuAuthority{})
}

func (i *initMenuAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil // do nothing
}

func (i *initMenuAuthority) TableCreated(ctx context.Context) bool {
	return false // always replace
}

func (i initMenuAuthority) InitializerName() string {
	return "vo1_role_mtm_menu"
}

func (i *initMenuAuthority) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	authorities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.Vo1Role)
	if !ok {
		return ctx, errors.Wrap(system.ErrMissingDependentContext, "创建 [菜单-权限] 关联失败, 未找到权限表初始化数据")
	}
	menus, ok := ctx.Value(initMenu{}.InitializerName()).([]sysModel.Vo1Menu)
	if !ok {
		return next, errors.Wrap(errors.New(""), "创建 [菜单-权限] 关联失败, 未找到菜单表初始化数据")
	}
	next = ctx
	// 888
	if err = db.Model(&authorities[0]).Association("Vo1Menu").Replace(menus[:20]); err != nil {
		return next, err
	}
	if err = db.Model(&authorities[0]).Association("Vo1Menu").Append(menus[21:]); err != nil {
		return next, err
	}

	// 8881
	menu8881 := menus[:2]
	menu8881 = append(menu8881, menus[7])
	if err = db.Model(&authorities[1]).Association("Vo1Menu").Replace(menu8881); err != nil {
		return next, err
	}

	// // 9528
	// if err = db.Model(&authorities[2]).Association("Vo1Menu").Replace(menus[:12]); err != nil {
	// 	return next, err
	// }
	// if err = db.Model(&authorities[2]).Association("Vo1Menu").Append(menus[13:17]); err != nil {
	// 	return next, err
	// }
	return next, nil
}

func (i *initMenuAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var count int64
	if err := db.Model(&sysModel.Vo1Role{}).
		Where("id = ?", "888").Preload("Vo1Menu").Count(&count); err != nil {
		return count == 16
	}
	return false
}
