package initialize

import (
	"github.com/zhangrt/voyager1_platform/config"
	"github.com/zhangrt/voyager1_platform/global"

	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.GS_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		case "pgsql":
			dbMap[info.AliasName] = GormPgSqlByConfig(config.Pgsql{GeneralDB: info.GeneralDB})
		case "cockroach":
			dbMap[info.AliasName] = GormCrockDbByConfig(config.Pgsql{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}

	if sysDB, ok := dbMap[sys]; ok {
		global.GS_DB = sysDB
	}
	global.GS_DBList = dbMap
}
