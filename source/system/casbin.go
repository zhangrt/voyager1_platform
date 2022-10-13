package system

import (
	"context"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/service/system"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const initOrderCasbin = system.InitOrderSystem + 1

type initCasbin struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCasbin, &initCasbin{})
}

func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
}

func (i *initCasbin) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&adapter.CasbinRule{})
}

func (i initCasbin) InitializerName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (i *initCasbin) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []adapter.CasbinRule{
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/base/login", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/admin_register", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authority/copyAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authority/updateAuthority", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authority/setDataAuthority", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/menu/getBaseMenuById", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/setUserInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/setSelfInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/deleteUser", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/setUserAuthorities", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/user/resetPassword", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/findFile", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/breakpointContinueFinish", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/breakpointContinue", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/removeChunk", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/uploadFile", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/getFileList", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/casbin/getPolicyPathByAuthorityId", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/jwt/jsonInBlacklist", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/system/getServerInfo", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/customer/customerList", V2: "GET"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/getDB", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/getMeta", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/preview", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/getTables", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/getColumn", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/rollback", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/createTemp", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/delSysHistory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/getSysHistory", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/createPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/getPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/delPackage", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/autoCode/createPlug", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionary/findSysDictionary", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionary/updateSysDictionary", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionary/getSysDictionaryList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionary/createSysDictionary", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysDictionary/deleteSysDictionary", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/email/emailTest", V2: "POST"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/simpleUploader/upload", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/simpleUploader/checkFileMd5", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/simpleUploader/mergeFileMd5", V2: "GET"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/excel/importExcel", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/excel/loadExcel", V2: "GET"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/excel/exportExcel", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/excel/downloadTemplate", V2: "GET"},

		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authorityBtn/setAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authorityBtn/getAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/authorityBtn/canRemoveAuthorityBtn", V2: "POST"},

		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/base/login", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/user/admin_register", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/api/createApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/api/getApiList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/api/getApiById", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/api/deleteApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/api/updateApi", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/api/getAllApis", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/authority/setDataAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/upload", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/getFileList", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/customer/customerList", V2: "GET"},
		{Ptype: "p", V0: "8881", V1: global.GS_CONFIG.System.Application + "/user/getUserInfo", V2: "GET"},

		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/base/login", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/user/admin_register", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/authority/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/authority/deleteAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/authority/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/authority/setDataAuthority", V2: "POST"},

		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/getBaseMenuTree", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/upload", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/getFileList", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/fileUploadAndDownload/editFileName", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/system/getSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/system/setSystemConfig", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "PUT"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/customer/customer", V2: "DELETE"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/customer/customerList", V2: "GET"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/autoCode/createTemp", V2: "POST"},
		{Ptype: "p", V0: "9528", V1: global.GS_CONFIG.System.Application + "/user/getUserInfo", V2: "GET"},
	}
	if err := db.Create(&entities).Error; err != nil {
		global.GS_LOG.Error("Casbin 表 ("+i.InitializerName()+") 数据初始化失败!", zap.Error(err))
		// return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitializerName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "888", V1: global.GS_CONFIG.System.Application + "/base/login", V2: "POST"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
