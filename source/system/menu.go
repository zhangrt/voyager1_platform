package system

import (
	"context"

	. "github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/service/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenu = initOrderAuthority + 1

type initMenu struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i initMenu) InitializerName() string {
	return Vo1Menu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&Vo1Menu{},
	)
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&Vo1Menu{})
}

func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []Vo1Menu{
		/*{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "1"}, MenuLevel: 0, Hidden: false, ParentId: "0", Url: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", SerialNo: 1, Description: "仪表盘", Icon: "odometer"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "2"}, MenuLevel: 0, Hidden: false, ParentId: "0", Url: "about", Name: "about", Component: "view/about/index.vue", SerialNo: 7, Description: "关于我们", Icon: "info-filled"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "3"}, MenuLevel: 0, Hidden: false, ParentId: "0", Url: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", SerialNo: 3, Description: "超级管理员", Icon: "user"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "4"}, MenuLevel: 0, Hidden: false, ParentId: "3", Url: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", SerialNo: 1, Description: "角色管理", Icon: "avatar"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "5"}, MenuLevel: 0, Hidden: false, ParentId: "3", Url: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", SerialNo: 2, Description: "菜单管理", Icon: "tickets"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "6"}, MenuLevel: 0, Hidden: false, ParentId: "3", Url: "api", Name: "api", Component: "view/superAdmin/api/api.vue", SerialNo: 3, Description: "api管理", Icon: "platform"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "7"}, MenuLevel: 0, Hidden: false, ParentId: "3", Url: "user", Name: "user", Component: "view/superAdmin/user/user.vue", SerialNo: 4, Description: "用户管理", Icon: "coordinate"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "8"}, MenuLevel: 0, Hidden: true, ParentId: "0", Url: "person", Name: "person", Component: "view/person/person.vue", SerialNo: 4, Description: "个人信息", Icon: "message"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "9"}, MenuLevel: 0, Hidden: false, ParentId: "0", Url: "example", Name: "example", Component: "view/example/index.vue", SerialNo: 6, Description: "示例文件", Icon: "management"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "10"}, MenuLevel: 0, Hidden: false, ParentId: "9", Url: "excel", Name: "excel", Component: "view/example/excel/excel.vue", SerialNo: 4, Description: "excel导入导出", Icon: "takeaway-box"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "11"}, MenuLevel: 0, Hidden: false, ParentId: "9", Url: "upload", Name: "upload", Component: "view/example/upload/upload.vue", SerialNo: 5, Description: "媒体库（上传下载）", Icon: "upload"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "12"}, MenuLevel: 0, Hidden: false, ParentId: "9", Url: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", SerialNo: 6, Description: "断点续传", Icon: "upload-filled"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "13"}, MenuLevel: 0, Hidden: false, ParentId: "9", Url: "customer", Name: "customer", Component: "view/example/customer/customer.vue", SerialNo: 7, Description: "客户列表（资源示例）", Icon: "avatar"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "14"}, MenuLevel: 0, Hidden: false, ParentId: "0", Url: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", SerialNo: 5, Description: "系统工具", Icon: "tools"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "15"}, MenuLevel: 0, Hidden: false, ParentId: "14", Url: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", SerialNo: 1, Description: "代码生成器", Icon: "cpu"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "16"}, MenuLevel: 0, Hidden: false, ParentId: "14", Url: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", SerialNo: 2, Description: "表单生成器", Icon: "magic-stick"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "17"}, MenuLevel: 0, Hidden: false, ParentId: "14", Url: "system", Name: "system", Component: "view/systemTools/system/system.vue", SerialNo: 3, Description: "系统配置", Icon: "operation"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "18"}, MenuLevel: 0, Hidden: false, ParentId: "3", Url: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", SerialNo: 5, Description: "字典管理", Icon: "notebook"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "19"}, MenuLevel: 0, Hidden: true, ParentId: "3", Url: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", SerialNo: 1, Description: "字典详情", Icon: "order"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "20"}, MenuLevel: 0, Hidden: false, ParentId: "3", Url: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", SerialNo: 6, Description: "操作历史", Icon: "pie-chart"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "21"}, MenuLevel: 0, Hidden: false, ParentId: "9", Url: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", SerialNo: 6, Description: "断点续传（插件版）", Icon: "upload"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "22"}, MenuLevel: 0, Hidden: false, ParentId: "0", Url: "http://www.github.com/zhangrt/voyager1_platform.com", Name: "http://www.github.com/zhangrt/voyager1_platform.com", Component: "/", SerialNo: 0, Description: "官方网站", Icon: "home-filled"},
		{GS_BASE_MODEL_ID_STRING: global.GS_BASE_MODEL_ID_STRING{ID: "23"}, MenuLevel: 0, Hidden: false, ParentId: "0", Url: "state", Name: "state", Component: "view/system/state.vue", SerialNo: 6, Description: "服务器状态", Icon: "cloudy"},*/
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, Vo1Menu{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("url = ?", "dashboard").First(&Vo1Menu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
