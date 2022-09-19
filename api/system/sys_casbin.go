package system

import (
	systemRes "github.com/zhangrt/voyager1_platform/model/system/response"
	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/response"

	"github.com/gin-gonic/gin"
	auth "github.com/zhangrt/voyager1_core/auth/luna"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body auth.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{msg=string} "更新角色api权限"
// @Router /casbin/UpdateCasbin [post]
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var req auth.CasbinInReceive
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	casbin := auth.NewCasbin()
	if err := casbin.UpdateCasbin(req.AuthorityId, req.CasbinInfos); err != nil {
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
// @Param data body auth.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{data=systemRes.PolicyPathResponse,msg=string} "获取权限列表,返回包括casbin详情列表"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *gin.Context) {
	var req auth.CasbinInReceive
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	casbin := auth.NewCasbin()
	paths := casbin.GetPolicyPathByAuthorityId(req.AuthorityId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
