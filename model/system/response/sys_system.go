package response

import "github.com/zhangrt/voyager1_platform/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
