package voyager1

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system"
)

type SystemApi struct{}

// @Tags Vo1System
// @Summary 根据子系统Id获取子系统信息
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetId true "子系统Id"
// @Success 200 {object} response.Response{msg=string} "子系统信息"
// @Router /system/getSystemModel [post]
func (s *SystemApi) GetSystemModel(id string) (_ system.Vo1System, err error) {
	var system system.Vo1System
	err = global.GS_DB.Where("global.GS_BASE_MODEL_ID_STRING.ID = ? and deleted <> 1", id).Error
	if err != nil {
		return *&system, err
	}
	return system, err
}

// @Tags Vo1System
// @Summary 新增或更新子系统
// @accept application/json
// @Produce application/json
// @Param data body systemReq.CreateSystemModel true "子系统模型"
// @Success 200 {object} response.Response{msg=string} "子系统模型"
// @Router /system/createSystemModel [post]
func (s *SystemApi) CreateSystemModel(systemReq system.Vo1System) (_ system.Vo1System, err error) {
	var system system.Vo1System
	err = global.GS_DB.Save(system).Error
	if err != nil {
		return *&system, err
	}
	return system, err
}
