package voyager1

import (
	systemRes "github.com/zhangrt/voyager1_platform/model/system/response"
	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/response"

	"github.com/gin-gonic/gin"
	"github.com/zhangrt/voyager1_core/auth/luna"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luna.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{msg=string} "更新角色api权限"
// @Router /v1/voyager/auth/updateCasbin [post]
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var req luna.CasbinInReceive
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	casbin := luna.NewCasbin()
	if err := casbin.UpdateCasbin(req.RoleId, req.CasbinInfos); err != nil {
		global.GS_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body luna.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{data=systemRes.PolicyPathResponse,msg=string} "获取权限列表,返回包括casbin详情列表"
// @Router /v1/voyager1/auth/getPolicyPathByAuthorityId [post]
func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *gin.Context) {
	var req luna.CasbinInReceive
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	casbin := luna.NewCasbin()
	paths := casbin.GetPolicyPathByAuthorityId(req.RoleId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
