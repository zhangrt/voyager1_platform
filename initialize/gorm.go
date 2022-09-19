package initialize

import (
	"github.com/zhangrt/voyager1_platform/global"
	DemoModels "github.com/zhangrt/voyager1_platform/model/demo"
	SystemModels "github.com/zhangrt/voyager1_platform/model/system"
	TestModels "github.com/zhangrt/voyager1_platform/model/test"

	auth "github.com/zhangrt/voyager1_core/auth/luna"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	global.GS_LOG.Info("gorm initialize process start")
	GS_DB := func() *gorm.DB {
		switch global.GS_CONFIG.System.DbType {
		case "mysql":
			return GormMysql()
		case "pgsql":
			return GormPgSql()
		case "cockroach":
			return GormCockroachDB()
		default:
			return GormCockroachDB()
		}
	}()
	global.GS_DB = GS_DB
	global.GS_LOG.Info("GS_DB:", zap.String("db", GS_DB.Name()))
	global.GS_LOG.Info("gorm initialize success")
	return GS_DB
}

// 自动建表,服务初始化时自动建表
func AutoMigrate() {
	var db = global.GS_DB
	// 实体
	var models = []interface{}{

		&TestModels.Test{},

		&DemoModels.Facility{},

		&SystemModels.SysUser{},
		&SystemModels.SysUseAuthority{},
		&SystemModels.SysAuthority{},
		&SystemModels.SysAuthorityBtn{},
		&SystemModels.SysBaseMenu{},
		&SystemModels.SysBaseMenuParameter{},
		&SystemModels.SysBaseMenuBtn{},
		&SystemModels.SysMenu{},
		&SystemModels.SysDictionary{},
		&SystemModels.SysDictionaryDetail{},
		&SystemModels.SysOperationRecord{},

		&auth.JwtBlacklist{},

		gormadapter.CasbinRule{},
	}
	// db.AutoMigrate(models...)
	for index := range models {
		if db.Migrator().HasTable(models[index]) {
			continue
		} else {
			db.Migrator().CreateTable(models[index])
		}
	}
}
