package system

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/zhangrt/voyager1_platform/config"

	"github.com/gookit/color"

	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system/request"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CockroachInitHandler struct{}

func NewCockroachInitHandler() *CockroachInitHandler {
	return &CockroachInitHandler{}
}

// WriteConfig Cockroach 回写配置
func (h CockroachInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Cockroach)
	if !ok {
		return errors.New("cockroach config invalid")
	}
	global.GS_CONFIG.System.DbType = "cockroach"
	global.GS_CONFIG.Cockroach = c
	global.GS_CONFIG.JWT.SigningKey = uuid.NewV4().String()
	cs := utils.StructToMap(global.GS_CONFIG)
	for k, v := range cs {
		global.GS_VP.Set(k, v)
	}
	return global.GS_VP.WriteConfig()
}

// EnsureDB 创建数据库并初始化 pg
func (h CockroachInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "cockroach" {
		return ctx, ErrDBTypeMismatch
	}
	dsn := conf.CockroachEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE %s;", conf.DBName)
	if err = createDatabase(dsn, "pgx", createSql); err != nil {
		return nil, err
	} // 创建数据库

	c := conf.ToCockroachConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	} // 如果没有数据库名, 则跳出初始化数据
	var db *gorm.DB
	if db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  c.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	global.GS_CONFIG.System.RootPath, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h CockroachInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h CockroachInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for i := 0; i < len(inits); i++ {
		if inits[i].DataInserted(next) {
			color.Info.Printf(InitDataExist, Cockroach, inits[i].InitializerName())
			continue
		}
		if n, err := inits[i].InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Cockroach, inits[i].InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Cockroach, inits[i].InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Cockroach)
	return nil
}
