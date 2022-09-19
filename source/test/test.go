package test

import (
	"context"
	testModel "github.com/zhangrt/voyager1_platform/model/test"
	"github.com/zhangrt/voyager1_platform/service/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderTest = system.InitOrderInternal + 1

type initTest struct{}

// auto run
func init() {
	system.RegisterInit(initOrderTest, &initTest{})
}

func (i *initTest) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&testModel.Test{})
}

func (i *initTest) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&testModel.Test{})
}

func (i initTest) InitializerName() string {
	return testModel.Test{}.TableName()
}

func (i *initTest) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []testModel.Test{
		{TestName: "初始化测试数据"},
	}

	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, testModel.Test{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initTest) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("test_name = ?", "初始化测试数据").First(&testModel.Test{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
