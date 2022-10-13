package system

import (
	"context"

	"github.com/zhangrt/voyager1_core/global"
	sysModel "github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/service/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderAuthority = initOrderCasbin + 1

type initAuthority struct{}

// auto run
func init() {
	system.RegisterInit(initOrderAuthority, &initAuthority{})
}

func (i *initAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.Vo1Role{})
}

func (i *initAuthority) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.Vo1Role{})
}

func (i initAuthority) InitializerName() string {
	return sysModel.Vo1Role{}.TableName()
}

func (i *initAuthority) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.Vo1Role{
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "888"}, Name: "普通用户"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "9528"}, Name: "测试角色"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "8881"}, Name: "普通用户子角色"},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!", sysModel.Vo1Role{}.TableName())
	}
	// data authority
	// if err := db.Model(&entities[0]).Association("DataAuthorityId").Replace(
	// 	[]*sysModel.Vo1Role{
	// 		{AuthorityId: "888"},
	// 		{AuthorityId: "9528"},
	// 		{AuthorityId: "8881"},
	// 	}); err != nil {
	// 	return ctx, errors.Wrapf(err, "%s表数据初始化失败!",
	// 		db.Model(&entities[0]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	// }
	// if err := db.Model(&entities[1]).Association("DataAuthorityId").Replace(
	// 	[]*sysModel.Vo1Role{
	// 		{AuthorityId: "9528"},
	// 		{AuthorityId: "8881"},
	// 	}); err != nil {
	// 	return ctx, errors.Wrapf(err, "%s表数据初始化失败!",
	// 		db.Model(&entities[1]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	// }

	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("id = ?", "888").
		First(&sysModel.Vo1Role{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
