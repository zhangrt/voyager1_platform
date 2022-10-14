package response

import "github.com/zhangrt/voyager1_platform/model/system"

type Vo1RoleResponse struct {
	Role system.Vo1Role `json:"role"`
}

type Vo1RoleCopyResponse struct {
	Role      system.Vo1Role `json:"role"`
	OldRoleId string         `json:"oldRoleId"` // 旧角色ID
}
