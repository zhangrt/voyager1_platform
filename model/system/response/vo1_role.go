package response

import "github.com/zhangrt/voyager1_platform/model/system"

type SysAuthorityResponse struct {
	Role system.Vo1Role `json:"role"`
}

type SysAuthorityCopyResponse struct {
	Role      system.Vo1Role `json:"role"`
	OldRoleId string         `json:"oldRoleId"` // 旧角色ID
}
