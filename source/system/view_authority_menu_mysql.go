package system

import (
	"context"
	"fmt"
	"strings"

	"github.com/zhangrt/voyager1_platform/service/system"

	"gorm.io/gorm"

	"github.com/zhangrt/voyager1_platform/global"
	sysModel "github.com/zhangrt/voyager1_platform/model/system"

	"github.com/pkg/errors"
)

const initOrderMenuViewMysql = initOrderMenuAuthority + 1

type initMenuViewMysql struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenuViewMysql, &initMenuViewMysql{})
}

func (i initMenuViewMysql) InitializerName() string {
	return fmt.Sprintf("mysql 视图<%s>", sysModel.Vo1Menu{}.TableName())
}

func (i *initMenuViewMysql) InitializeData(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (i *initMenuViewMysql) DataInserted(ctx context.Context) bool {
	return true // ignore
}

func (v *initMenuViewMysql) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" {
		return ctx, nil // ignore
	}
	joinTableName := db.Model(&sysModel.Vo1Role{}).Association("SysBaseMenus").Relationship.JoinTable.Name
	sql := `
	CREATE OR REPLACE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW @table_name AS
	select @menus.id                AS id,
		   @menus.path              AS path,
		   @menus.icon              AS icon,
		   @menus.name              AS name,
		   @menus.sort              AS sort,
		   @menus.title             AS title,
		   @menus.hidden            AS hidden,
		   @menus.component         AS component,
		   @menus.parent_id         AS parent_id,
		   @menus.created_at        AS created_at,
		   @menus.updated_at        AS updated_at,
		   @menus.deleted_at        AS deleted_at,
		   @menus.keep_alive        AS keep_alive,
		   @menus.menu_level        AS menu_level,
		   @menus.default_menu      AS default_menu,
		   @menus.close_tab      	AS close_tab,
		   @role_menus.id      AS menu_id,
		   @role_menus.role_id AS role_id
	from (@role_menus
			 join @menus on ((@role_menus.menu_id = @menus.id)));
	`
	sql = strings.ReplaceAll(sql, "@table_name", sysModel.Vo1Menu{}.TableName())
	sql = strings.ReplaceAll(sql, "@menus", sysModel.Vo1Menu{}.TableName())
	sql = strings.ReplaceAll(sql, "@role_menus", joinTableName)
	if err := db.Exec(sql).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.Vo1Menu{}.TableName()+"视图创建失败!")
	}
	return ctx, nil
}

func (i *initMenuViewMysql) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	err1 := db.Find(&[]sysModel.Vo1Menu{}).Error
	err2 := errors.New(fmt.Sprintf("Error 1146: Table '%v.%v' doesn't exist",
		global.GS_CONFIG.Mysql.Dbname, sysModel.Vo1Menu{}.TableName()))
	if errors.As(err1, &err2) {
		return false
	}
	return true
}
