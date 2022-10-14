package request

import (
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/system"
)

type Vo1DictionarySearch struct {
	system.Vo1Dictionary
	request.PageInfo
}
