package system

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/response"
	"github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/model/system/request"
	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryApi struct{}

// @Tags Vo1Dictionary
// @Summary 创建Vo1Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Dictionary true "Vo1Dictionary模型"
// @Success 200 {object} response.Response{msg=string} "创建Vo1Dictionary"
// @Router /dictionary/createDictionary [post]
func (s *DictionaryApi) CreateVo1Dictionary(c *gin.Context) {
	var dictionary system.Vo1Dictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.CreateVo1Dictionary(dictionary); err != nil {
		global.GS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Vo1Dictionary
// @Summary 删除Vo1Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Dictionary true "Vo1Dictionary模型"
// @Success 200 {object} response.Response{msg=string} "删除Vo1Dictionary"
// @Router /dctionary/deleteDictionary [delete]
func (s *DictionaryApi) DeleteVo1Dictionary(c *gin.Context) {
	var dictionary system.Vo1Dictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.DeleteVo1Dictionary(dictionary); err != nil {
		global.GS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Vo1Dictionary
// @Summary 更新Vo1Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1Dictionary true "Vo1Dictionary模型"
// @Success 200 {object} response.Response{msg=string} "更新Vo1Dictionary"
// @Router /dictionary/updateDictionary [put]
func (s *DictionaryApi) UpdateVo1Dictionary(c *gin.Context) {
	var dictionary system.Vo1Dictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.UpdateVo1Dictionary(&dictionary); err != nil {
		global.GS_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Vo1Dictionary
// @Summary 用id查询Vo1Dictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.Vo1Dictionary true "ID或字典英名"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询Vo1Dictionary"
// @Router /dictionary/findDictionary [get]
func (s *DictionaryApi) FindVo1Dictionary(c *gin.Context) {
	var dictionary system.Vo1Dictionary
	_ = c.ShouldBindQuery(&dictionary)
	if Vo1Dictionary, err := dictionaryService.GetVo1Dictionary(dictionary.Type, dictionary.ID); err != nil {
		global.GS_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reVo1Dictionary": Vo1Dictionary}, "查询成功", c)
	}
}

// @Tags Vo1Dictionary
// @Summary 分页获取Vo1Dictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.Vo1DictionarySearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取Vo1Dictionary列表,返回包括列表,总数,页码,每页数量"
// @Router /dictionary/getDictionaryList [get]
func (s *DictionaryApi) GetVo1DictionaryList(c *gin.Context) {
	var pageInfo request.Vo1DictionarySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dictionaryService.GetVo1DictionaryInfoList(pageInfo); err != nil {
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
