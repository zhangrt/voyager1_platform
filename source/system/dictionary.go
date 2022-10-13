package system

import (
	"context"

	"github.com/zhangrt/voyager1_core/global"
	sysModel "github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/service/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderDict = initOrderCasbin + 1

type initDict struct{}

// auto run
func init() {
	system.RegisterInit(initOrderDict, &initDict{})
}

func (i *initDict) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.Vo1Dictionary{})
}

func (i *initDict) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.Vo1Dictionary{})
}

func (i initDict) InitializerName() string {
	return sysModel.Vo1Dictionary{}.TableName()
}

func (i *initDict) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.Vo1Dictionary{
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "gender"}, NameCN: "性别", NameEN: "gender", Description: "性别字典"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "int"}, NameCN: "数据库int类型", NameEN: "int", Description: "int类型对应的数据库类型"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "time"}, NameCN: "数据库时间日期类型", NameEN: "time.Time", Description: "数据库时间日期类型"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "float64"}, NameCN: "数据库浮点型", NameEN: "float64", Description: "数据库浮点型"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "string"}, NameCN: "数据库字符串", NameEN: "string", Description: "数据库字符串"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "bool"}, NameCN: "数据库bool类型", NameEN: "bool", Description: "数据库bool类型"},
	}

	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.Vo1Dictionary{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initDict) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("id = ?", "gender").First(&sysModel.Vo1Dictionary{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
