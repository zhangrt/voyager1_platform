package test

import (
	"context"
	"encoding/json"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/test"
	"github.com/zhangrt/voyager1_platform/model/test/request"

	"github.com/zhangrt/voyager1_core/cache"
)

type TestRepository struct{}

func (tTestRepository *TestRepository) TestGet(ctx context.Context, id string) (data string, err error) {
	get_back := cache.Get(id)
	return get_back, err
}

func (tTestRepository *TestRepository) TestUpdate(ctx context.Context, testName string, ids ...string) {
	originSql := "update test set testName=? where user_id in (?) "
	global.GS_DB.Exec(originSql, testName, ids)
	return
}

func (tTestRepository *TestRepository) TestSelect(ctx context.Context, id string) (testName string, err error) {
	originSql := "select testName from test where id = ? "
	rows, err := global.GS_DB.Raw(originSql, id).Rows()
	defer rows.Close()
	if rows != nil {
		for rows.Next() {
			rows.Scan(&testName)
		}
	}
	return testName, err
}

func (tTestRepository *TestRepository) GetTestInfoList(info request.TestSearch) (list []test.Test, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GS_DB.Model(&test.Test{})
	json, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		global.GS_LOG.Error(err.Error())
	}
	global.GS_LOG.Info(string(json))
	var testList []test.Test
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TestName != "" {
		db = db.Where("test_name LIKE ?", "%"+info.TestName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&testList).Error
	return testList, total, err
}
