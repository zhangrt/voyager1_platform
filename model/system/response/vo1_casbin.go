package response

import auth "github.com/zhangrt/voyager1_core/auth/luna"

type PolicyPathResponse struct {
	Paths []auth.CasbinInfo `json:"paths"`
}
