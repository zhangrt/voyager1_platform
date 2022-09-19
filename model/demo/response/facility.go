package response

import (
	"github.com/zhangrt/voyager1_platform/model/demo"
)

type FacilityResponse struct {
	Type         string          `json:"type"`
	FacilityList []demo.Facility `json:"facilityList"`
}
