package system

import (
	"strconv"

	core "github.com/zhangrt/voyager1_core/global"
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/common/response"
	"github.com/zhangrt/voyager1_platform/model/system"
	systemReq "github.com/zhangrt/voyager1_platform/model/system/request"
	systemRes "github.com/zhangrt/voyager1_platform/model/system/response"
	"github.com/zhangrt/voyager1_platform/utils"

	auth "github.com/zhangrt/voyager1_core/auth/luna"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// var store = base64Captcha.DefaultMemStore

type UserApi struct{}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/login [post]
func (b *UserApi) Login(c *gin.Context) {
	var l systemReq.Login
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// if store.Verify(l.CaptchaId, l.Captcha, true) {
	u := &system.Vo1Person{
		GS_BASE_USER: core.GS_BASE_USER{
			Account:  l.Identification,
			Phone:    l.Identification,
			Email:    l.Identification,
			Password: l.Password,
		},
		OrganizationId: l.OrganizationId,
	}
	if user, err := userService.Login(u); err != nil {
		global.GS_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		b.tokenNext(c, *user)
	}
	// } else {
	// 	response.FailWithMessage("验证码错误", c)
	// }
}

// 登录以后签发jwt
func (b *UserApi) tokenNext(c *gin.Context, user system.Vo1Person) {
	j := &auth.TOKEN{SigningKey: []byte(global.GS_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(auth.BaseClaims{
		ID:              user.ID,
		Name:            user.Name,
		Account:         user.Account,
		Phone:           user.Phone,
		Email:           user.Email,
		RoleIds:         user.RoleIds,
		Roles:           user.Roles,
		DepartMentIds:   user.DepartmentIds,
		DepartMents:     user.Departments,
		OrganizationId:  user.OrganizationId,
		OrganizationIds: user.OrganizationIds,
		Organizations:   user.Organizations,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GS_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GS_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	var jwtService = auth.NewJWT()

	if jwtStr, err := jwtService.GetCacheJWT(user.Account); err == redis.Nil {
		if err := jwtService.SetCacheJWT(token, user.Account); err != nil {
			global.GS_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GS_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT auth.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetCacheJWT(token, user.Account); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// @Tags Vo1Person
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=systemRes.Vo1PersonResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/admin_register [post]
func (b *UserApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.Vo1Role
	for _, v := range r.RoleIds {
		authorities = append(authorities, system.Vo1Role{
			GS_BASE_MODEL_ID_STRING: core.GS_BASE_MODEL_ID_STRING{
				ID: v,
			},
		})
	}

	user := &system.Vo1Person{
		Roles: authorities,
	}
	user.Account = r.Account
	user.Name = r.Name
	user.Password = r.Password
	user.Avatar = r.Avatar
	user.RoleIds = r.RoleIds
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.GS_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.Vo1PersonResponse{Person: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.Vo1PersonResponse{Person: userReturn}, "注册成功", c)
	}
}

// @Tags Vo1Person
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {object} response.Response{msg=string} "用户修改密码"
// @Router /user/changePassword [post]
func (b *UserApi) ChangePassword(c *gin.Context) {
	var user systemReq.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.Vo1Person{}
	u.Account = user.Account
	u.Password = user.Password
	if _, err := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.GS_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags Vo1Person
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router /user/getUserList [post]
func (b *UserApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userService.GetUserInfoList(pageInfo); err != nil {
		global.GS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Vo1Person
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuthorities true "用户UUID, 角色ID"
// @Success 200 {object} response.Response{msg=string} "设置用户权限"
// @Router /user/setUserAuthorities [post]
func (b *UserApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	_ = c.ShouldBindJSON(&sua)
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	if err := userService.SetUserAuthorities(sua.ID, sua.RoleIds); err != nil {
		global.GS_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		claims := auth.GetUserInfo(c)
		j := &auth.TOKEN{SigningKey: []byte(global.GS_CONFIG.JWT.SigningKey)} // 唯一签名
		claims.RoleIds = sua.RoleIds
		if token, err := j.CreateToken(*claims); err != nil {
			global.GS_LOG.Error("修改失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
		} else {
			c.Header("new-token", token)
			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			response.OkWithMessage("修改成功", c)
		}

		response.OkWithMessage("修改成功", c)
	}
}

// @Tags Vo1Person
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {object} response.Response{msg=string} "删除用户"
// @Router /user/deleteUser [delete]
func (b *UserApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := auth.GetUserID(c)
	if jwtId == reqId.ID {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		global.GS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Vo1Person
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Person true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/setUserInfo [put]
func (b *UserApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(user.RoleIds) != 0 {
		err := userService.SetUserAuthorities(user.ID.String(), user.RoleIds)
		if err != nil {
			global.GS_LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", c)
		}
	}

	if err := userService.SetUserInfo(system.Vo1Person{
		GS_BASE_USER: core.GS_BASE_USER{
			ID:     user.ID,
			Name:   user.Name,
			Avatar: user.Avatar,
			Phone:  user.Phone,
			Email:  user.Email,
		},
	}); err != nil {
		global.GS_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// @Tags Vo1Person
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Person true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/SetSelfInfo [put]
func (b *UserApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	_ = c.ShouldBindJSON(&user)
	user.ID = auth.GetUserUUID(c)
	if err := userService.SetUserInfo(system.Vo1Person{
		GS_BASE_USER: core.GS_BASE_USER{
			ID:     user.ID,
			Name:   user.Name,
			Avatar: user.Avatar,
			Phone:  user.Phone,
			Email:  user.Email,
		},
	}); err != nil {
		global.GS_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// @Tags Vo1Person
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取用户信息"
// @Router /user/getUserInfo [get]
func (b *UserApi) GetUserInfo(c *gin.Context) {
	uuid := auth.GetUserUUID(c)
	if ReqUser, err := userService.GetUserInfo(uuid); err != nil {
		global.GS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

// @Tags Vo1Person
// @Summary 重置用户密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.Vo1Person true "ID"
// @Success 200 {object} response.Response{msg=string} "重置用户密码"
// @Router /user/resetPassword [post]
func (b *UserApi) ResetPassword(c *gin.Context) {
	var user system.Vo1Person
	_ = c.ShouldBindJSON(&user)
	if err := userService.ResetPassword(user.ID.String()); err != nil {
		global.GS_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("重置成功", c)
	}
}
