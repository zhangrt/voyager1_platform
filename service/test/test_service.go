package test

import (
	"context"
	"strconv"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/test"
	"github.com/zhangrt/voyager1_platform/model/test/request"

	redis "github.com/zhangrt/voyager1_core/cache"
)

type TestService struct {
	cache redis.Cacher
}

func (testService *TestService) TestPost(id uint, name string) (data string, err error) {
	// uint => string
	var test test.Test
	test.ID = id
	test.TestName = name
	err = global.GS_DB.Create(&test).Error
	if err == nil {
		_ = testService.cache.Set(strconv.Itoa(int(id)), name, 0)
	}
	if err != nil {
		return err.Error(), err
	}
	return "test post is OKay", err
}

func (testService *TestService) TestGet(id string) (data string, err error) {
	get_back, err := testRepository.TestGet(context.Background(), id)
	return get_back, err
}

func (testService *TestService) GetTestInfoList(info request.TestSearch) (list interface{}, total int64, err error) {
	var testList []test.Test
	testList, total, err = testRepository.GetTestInfoList(info)
	return testList, total, err
}
