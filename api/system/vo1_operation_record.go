package system

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/common/response"
	"github.com/zhangrt/voyager1_platform/model/system"
	systemReq "github.com/zhangrt/voyager1_platform/model/system/request"
	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationRecordApi struct{}

// @Tags Vo1OperationRecord
// @Summary 创建Vo1OperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1OperationRecord true "创建Vo1OperationRecord"
// @Success 200 {object} response.Response{msg=string} "创建Vo1OperationRecord"
// @Router /Vo1OperationRecord/createVo1OperationRecord [post]
func (s *OperationRecordApi) CreateVo1OperationRecord(c *gin.Context) {
	var Vo1OperationRecord system.Vo1OperationRecord
	_ = c.ShouldBindJSON(&Vo1OperationRecord)
	if err := operationRecordService.CreateVo1OperationRecord(Vo1OperationRecord); err != nil {
		global.GS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Vo1OperationRecord
// @Summary 删除Vo1OperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.Vo1OperationRecord true "Vo1OperationRecord模型"
// @Success 200 {object} response.Response{msg=string} "删除Vo1OperationRecord"
// @Router /Vo1OperationRecord/deleteVo1OperationRecord [delete]
func (s *OperationRecordApi) DeleteVo1OperationRecord(c *gin.Context) {
	var Vo1OperationRecord system.Vo1OperationRecord
	_ = c.ShouldBindJSON(&Vo1OperationRecord)
	if err := operationRecordService.DeleteVo1OperationRecord(Vo1OperationRecord); err != nil {
		global.GS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Vo1OperationRecord
// @Summary 批量删除Vo1OperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Vo1OperationRecord"
// @Success 200 {object} response.Response{msg=string} "批量删除Vo1OperationRecord"
// @Router /Vo1OperationRecord/deleteVo1OperationRecordByIds [delete]
func (s *OperationRecordApi) DeleteVo1OperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := operationRecordService.DeleteVo1OperationRecordByIds(IDS); err != nil {
		global.GS_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Vo1OperationRecord
// @Summary 用id查询Vo1OperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.Vo1OperationRecord true "Id"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询Vo1OperationRecord"
// @Router /Vo1OperationRecord/findVo1OperationRecord [get]
func (s *OperationRecordApi) FindVo1OperationRecord(c *gin.Context) {
	var Vo1OperationRecord system.Vo1OperationRecord
	_ = c.ShouldBindQuery(&Vo1OperationRecord)
	if err := utils.Verify(Vo1OperationRecord, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reVo1OperationRecord, err := operationRecordService.GetVo1OperationRecord(Vo1OperationRecord.ID); err != nil {
		global.GS_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reVo1OperationRecord": reVo1OperationRecord}, "查询成功", c)
	}
}

// @Tags Vo1OperationRecord
// @Summary 分页获取Vo1OperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.Vo1OperationRecordSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取Vo1OperationRecord列表,返回包括列表,总数,页码,每页数量"
// @Router /Vo1OperationRecord/getVo1OperationRecordList [get]
func (s *OperationRecordApi) GetVo1OperationRecordList(c *gin.Context) {
	var pageInfo systemReq.Vo1OperationRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := operationRecordService.GetVo1OperationRecordInfoList(pageInfo); err != nil {
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
