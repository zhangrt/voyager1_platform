package demo

import (
	"testing"

	r2 "github.com/zhangrt/voyager1_platform/model/common/request"
	r1 "github.com/zhangrt/voyager1_platform/model/demo/request"
)

var (
	repository FacilityRepository
)

func TestGetFacilityInfoList(t *testing.T) {

	info := r1.FacilitySearch{
		Name: "1",
		PageInfo: r2.PageInfo{
			PageSize: 10,
			Page:     1,
			Keyword:  "type-desc,status-desc",
		},
	}

	repository.GetFacilityInfoList(info)
}
