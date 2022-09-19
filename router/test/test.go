package test

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (e *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	testApi := v1.ApiGroupApp.TestApiGroup.TestApi
	{
		testRouter.GET("get", testApi.TestServer)           // test get
		testRouter.GET("server", testApi.TestMircoServer)   // test get
		testRouter.POST("testPost", testApi.TestPost)       // test post
		testRouter.GET("testGet", testApi.TestGet)          // test get
		testRouter.GET("getTestList", testApi.GetTestList)  // test get list
		testRouter.PUT("testPut", testApi.TestPut)          // test put
		testRouter.DELETE("testDelete", testApi.TestDelete) // test delete
	}
}
