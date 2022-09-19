package demo

import (
	"encoding/json"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/response"
	"github.com/zhangrt/voyager1_platform/model/demo"
	"github.com/zhangrt/voyager1_platform/model/demo/request"
	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FacilityApi struct{}

// @Tags FacilityApi
// @Summary 添加设备
// @accept application/json
// @Produce application/json
// @Param data body request.FacilityAdd true "添加设备"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /facility/add [post]
func (s *FacilityApi) AddFacility(c *gin.Context) {
	var info request.FacilityAdd
	_ = c.ShouldBindJSON(&info)

	json, err := json.MarshalIndent(info, "", " ")
	global.GS_LOG.Info(string(json), zap.Error(err))

	if err := utils.Verify(info, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if data, err := facilityService.AddFacility(request.AddToFacility(info)); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(data, c)
	}

}

// @Tags FacilityApi
// @Summary 删除设备
// @accept application/json
// @Produce application/json
// @Param id query uint true "通过设备id删除设备"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /facility/remove [delete]
func (s *FacilityApi) RemoveFacility(c *gin.Context) {
	var info demo.Facility
	_ = c.ShouldBindQuery(&info)

	if data, err := facilityService.RemoveFacility(info); err != nil {
		global.GS_LOG.Error("remove facility failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(data, c)
	}

}

// @Tags FacilityApi
// @Summary 更新设备
// @accept application/json
// @Produce application/json
// @Param data body request.FacilityUpdate true "更新设备"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /facility/update [put]
func (s *FacilityApi) UpdateFacility(c *gin.Context) {
	var info request.FacilityUpdate
	_ = c.ShouldBindJSON(&info)

	json, err := json.MarshalIndent(info, "", " ")
	global.GS_LOG.Info(string(json), zap.Error(err))

	if err := utils.Verify(info, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if data, err := facilityService.UpdateFacility(request.UpdateToFacility(info)); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(data, c)
	}

}

// @Tags FacilityApi
// @Summary 获取设备详情
// @Produce application/json
// @Param id query uint true "通过设备id获取设备详情"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /facility/get [get]
func (e *FacilityApi) GetFacility(c *gin.Context) {
	id := c.Query("id")

	if data, err := facilityService.GetFacility(id); err != nil {
		global.GS_LOG.Error("get failed by: "+id, zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(data, c)
	}
}

// @Tags FacilityApi
// @Summary 分页获取设备列表
// @accept application/json
// @Produce application/json
// @Param data query request.FacilitySearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取设备列表,返回包括列表,总数,页码,每页数量"
// @Router /facility/getList [get]
func (s *FacilityApi) GetFacilityList(c *gin.Context) {
	var pageInfo request.FacilitySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := facilityService.GetFacilityInfoList(pageInfo); err != nil {
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
