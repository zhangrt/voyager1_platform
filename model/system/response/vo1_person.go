package response

import (
	"github.com/zhangrt/voyager1_platform/model/system"
)

type Vo1PersonResponse struct {
	Person system.Vo1Person `json:"person"`
}

type LoginResponse struct {
	Person    system.Vo1Person `json:"person"`
	Token     string           `json:"token"`
	ExpiresAt int64            `json:"expiresAt"`
}
