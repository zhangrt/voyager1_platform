package request

import (
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/system"
)

type Vo1OperationRecordSearch struct {
	system.Vo1OperationRecord
	request.PageInfo
}
