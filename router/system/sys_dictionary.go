package system

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitDictionaryRouter(Router *gin.RouterGroup) {
	Vo1DictionaryRouter := Router.Group("dictionary")
	Vo1DictionaryRouterWithoutRecord := Router.Group("dictionary")
	Vo1DictionaryApi := v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		Vo1DictionaryRouter.POST("createDictionary", Vo1DictionaryApi.CreateVo1Dictionary)   // 新建Vo1Dictionary
		Vo1DictionaryRouter.DELETE("deleteDictionary", Vo1DictionaryApi.DeleteVo1Dictionary) // 删除Vo1Dictionary
		Vo1DictionaryRouter.PUT("updateDictionary", Vo1DictionaryApi.UpdateVo1Dictionary)    // 更新Vo1Dictionary
	}
	{
		Vo1DictionaryRouterWithoutRecord.GET("findDictionary", Vo1DictionaryApi.FindVo1Dictionary)       // 根据ID获取Vo1Dictionary
		Vo1DictionaryRouterWithoutRecord.GET("getDictionaryList", Vo1DictionaryApi.GetVo1DictionaryList) // 获取Vo1Dictionary列表
	}
}
