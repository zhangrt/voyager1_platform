package system

import (
	"context"

	sysModel "github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/service/system"
	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
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
	return ctx, db.AutoMigrate(&sysModel.SysUser{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	password := utils.BcryptHash("q123456.")
	adminPassword := utils.BcryptHash("q123456.")

	entities := []sysModel.SysUser{
		{
			UUID:        uuid.NewV4(),
			Username:    "admin",
			Password:    adminPassword,
			NickName:    "超级管理员",
			HeaderImg:   "https://img1.baidu.com/it/u=2838100141,2488760005&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1662483600&t=9f7622bd7a5ab6abda1ec1de3e8797af",
			AuthorityId: "888",
			Phone:       "17611111111",
			Email:       "333333333@qq.com",
		},
		{
			UUID:        uuid.NewV4(),
			Username:    "zhoujj",
			Password:    adminPassword,
			NickName:    "ZHOUJIAJUN",
			HeaderImg:   "https://img1.baidu.com/it/u=2838100141,2488760005&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1662483600&t=9f7622bd7a5ab6abda1ec1de3e8797af",
			AuthorityId: "888",
			Phone:       "18966668888",
			Email:       "zhoujiajun@github.com/zhangrt/voyager1_platform.com",
		},
		{
			UUID:        uuid.NewV4(),
			Username:    "test",
			Password:    password,
			NickName:    "BIG Monster",
			HeaderImg:   "https://img0.baidu.com/it/u=4060770951,4069855872&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1662483600&t=92a8aac26b4757fe849a8a10aaf31d87",
			AuthorityId: "9528",
			Phone:       "17611111111",
			Email:       "333333333@qq.com"},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	authorityEntities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	if err = db.Model(&entities[0]).Association("Authorities").Replace(authorityEntities); err != nil {
		return next, err
	}
	if err = db.Model(&entities[1]).Association("Authorities").Replace(authorityEntities[:1]); err != nil {
		return next, err
	}
	return next, err
}

func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysUser
	if errors.Is(db.Where("username = ?", "a303176530").
		Preload("Authorities").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Authorities) > 0 && record.Authorities[0].AuthorityId == "888"
}