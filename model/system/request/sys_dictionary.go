package request

import (
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
