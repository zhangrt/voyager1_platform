package voyager1

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/common/response"
	"github.com/zhangrt/voyager1_platform/model/system"
	systemReq "github.com/zhangrt/voyager1_platform/model/system/request"
	systemRes "github.com/zhangrt/voyager1_platform/model/system/response"
	"github.com/zhangrt/voyager1_platform/utils"

	auth "github.com/zhangrt/voyager1_core/auth/luna"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleApi struct{}

// @Tags Authority
// @Summary 通过用户角色信息获取菜单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.GetMenusByRoleIds true "权限ids"
// @Success 200 {object} response.Response{data=systemRes.Vo1MenusResponse,msg=string} "创建角色,返回包括系统角色详情"
// @Router /v1/voyager1/auth/getMenusByRoleIds [post]
func (a *RoleApi) GetMenusByRoleIds(c *gin.Context) {
	var req systemReq.GetMenusByRoleIds
	err := c.ShouldBindJSON(&req)
	if err != nil || req.RoleIds == nil || len(req.RoleIds) == 0 {
		req.RoleIds = auth.GetUserAuthorityId(c)
	}
	if err := utils.Verify(req, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if menus, err := roleService.GetMenusByRoleIds(req); err != nil {
		global.GS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(menus, "获取成功", c)
	}
}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Role true "权限id, 权限名, 父角色id"
// @Success 200 {object} response.Response{data=systemRes.Vo1RoleResponse,msg=string} "创建角色,返回包括系统角色详情"
// @Router /authority/createAuthority [post]
func (a *RoleApi) CreateAuthority(c *gin.Context) {
	var authority system.Vo1Role
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authBack, err := roleService.CreateAuthority(authority); err != nil {
		global.GS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.ID)
		casbin := auth.NewCasbin()
		_ = casbin.UpdateCasbin(authority.ID, auth.DefaultCasbin())
		response.OkWithDetailed(systemRes.Vo1RoleResponse{Role: authBack}, "创建成功", c)
	}
}

// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.Vo1RoleCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {object} response.Response{data=systemRes.Vo1RoleResponse,msg=string} "拷贝角色,返回包括系统角色详情"
// @Router /authority/copyAuthority [post]
func (a *RoleApi) CopyAuthority(c *gin.Context) {
	var copyInfo systemRes.Vo1RoleCopyResponse
	_ = c.ShouldBindJSON(&copyInfo)
	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(copyInfo.Role, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authBack, err := roleService.CopyAuthority(copyInfo); err != nil {
		global.GS_LOG.Error("拷贝失败!", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.Vo1RoleResponse{Role: authBack}, "拷贝成功", c)
	}
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Role true "删除角色"
// @Success 200 {object} response.Response{msg=string} "删除角色"
// @Router /authority/deleteAuthority [post]
func (a *RoleApi) DeleteAuthority(c *gin.Context) {
	var authority system.Vo1Role
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.GS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Role true "权限id, 权限名, 父角色id"
// @Success 200 {object} response.Response{data=systemRes.Vo1RoleResponse,msg=string} "更新角色信息,返回包括系统角色详情"
// @Router /authority/updateAuthority [post]
func (a *RoleApi) UpdateAuthority(c *gin.Context) {
	var auth system.Vo1Role
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authority, err := roleService.UpdateAuthority(auth); err != nil {
		global.GS_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.Vo1RoleResponse{Role: authority}, "更新成功", c)
	}
}

// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取角色列表,返回包括列表,总数,页码,每页数量"
// @Router /authority/getAuthorityList [post]
func (a *RoleApi) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := roleService.GetAuthorityInfoList(pageInfo); err != nil {
		global.GS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Role true "设置角色资源权限"
// @Success 200 {object} response.Response{msg=string} "设置角色资源权限"
// @Router /authority/setDataAuthority [post]
func (a *RoleApi) SetDataAuthority(c *gin.Context) {
	var auth system.Vo1Role
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.SetDataAuthority(auth); err != nil {
		global.GS_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}
