package system

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userRouterWithoutRecord := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("login", userApi.Login)
		userRouter.POST("admin_register", userApi.Register)               // 管理员注册账号
		userRouter.POST("changePassword", userApi.ChangePassword)         // 用户修改密码
		userRouter.DELETE("deleteUser", userApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", userApi.SetUserInfo)                // 设置用户信息
		userRouter.PUT("setSelfInfo", userApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", userApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", userApi.ResetPassword)           // 设置用户权限组
	}
	{
		userRouterWithoutRecord.POST("getUserList", userApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", userApi.GetUserInfo)  // 获取自身信息
	}
}
