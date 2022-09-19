package request

import (
	"github.com/zhangrt/voyager1_platform/model/test"

	"github.com/zhangrt/voyager1_platform/model/common/request"
)

type TestSearch struct {
	test.Test
	request.PageInfo
}
