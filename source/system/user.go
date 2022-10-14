package system

import (
	"context"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/global"
	sysModel "github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/service/system"
	"github.com/zhangrt/voyager1_platform/utils"
	"gorm.io/gorm"
)

const initOrderUser = initOrderAuthority + 1

type initUser struct{}

// auto run
func init() {
	system.RegisterInit(initOrderUser, &initUser{})
}

func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.Vo1Person{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.Vo1Person{})
}

func (i initUser) InitializerName() string {
	return sysModel.Vo1Person{}.TableName()
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	password := utils.BcryptHash("q123456.")
	adminPassword := utils.BcryptHash("q123456.")

	entities := []sysModel.Vo1Person{
		{
			Roles: []sysModel.Vo1Role{
				{
					GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{
						ID: "888",
					},
				},
			},
			GS_BASE_USER: global.GS_BASE_USER{
				UUID:           uuid.NewV4(),
				Account:        "admin",
				Password:       adminPassword,
				Name:           "超级管理员",
				Avatar:         "https://img1.baidu.com/it/u=2838100141,2488760005&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1662483600&t=9f7622bd7a5ab6abda1ec1de3e8797af",
				Phone:          "17611111111",
				Email:          "333333333@qq.com",
				DepartMentId:   "111",
				OrganizationId: "11",
			},
		},
		{
			Roles: []sysModel.Vo1Role{
				{
					GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{
						ID: "888",
					},
				},
			},
			GS_BASE_USER: global.GS_BASE_USER{
				UUID:           uuid.NewV4(),
				Account:        "zhoujj",
				Password:       adminPassword,
				Name:           "ZHOUJIAJUN",
				Avatar:         "https://img1.baidu.com/it/u=2838100141,2488760005&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1662483600&t=9f7622bd7a5ab6abda1ec1de3e8797af",
				Phone:          "18966668888",
				Email:          "zhoujiajun@github.com/zhangrt/voyager1_platform.com",
				DepartMentId:   "111",
				OrganizationId: "11",
			},
		},
		{
			Roles: []sysModel.Vo1Role{
				{
					GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{
						ID: "888",
					},
				},
			},
			GS_BASE_USER: global.GS_BASE_USER{
				UUID:           uuid.NewV4(),
				Account:        "test",
				Password:       password,
				Name:           "BIG Monster",
				Avatar:         "https://img0.baidu.com/it/u=4060770951,4069855872&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1662483600&t=92a8aac26b4757fe849a8a10aaf31d87",
				Phone:          "17611111111",
				Email:          "333333333@qq.com",
				DepartMentId:   "222",
				OrganizationId: "12",
			},
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.Vo1Person{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	authorityEntities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.Vo1Role)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("Roles").Replace(authorityEntities); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("Roles").Replace(authorityEntities[:1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.Vo1Person
	if errors.Is(db.Where("account = ?", "admin").
		Preload("Roles").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Roles) > 0 && record.Roles[0].ID == "888"
}
