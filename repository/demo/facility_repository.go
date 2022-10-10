package demo

import (
	"context"
	"strconv"
	"strings"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/demo"
	"github.com/zhangrt/voyager1_platform/model/demo/request"
	"github.com/zhangrt/voyager1_platform/utils"

	"go.uber.org/zap"
)

var (
	builder utils.SQLBuilder
)

type FacilityRepository struct{}

func (facilityRepository *FacilityRepository) GetFacility(ctx context.Context, id string) (data demo.Facility, err error) {
	var facility demo.Facility
	param, err := strconv.Atoi(id)
	if err != nil {
		return facility, err
	}
	err = global.GS_DB.Where("id", param).First(&facility).Error
	return facility, err
}

func (facilityRepository *FacilityRepository) GetFacilityInfoList(info request.FacilitySearch) (list []demo.Facility, total int64, err error) {
	var facilityList []demo.Facility
	if global.GS_DB == nil {
		return facilityList, total, err
	}

	// // 创建db
	// db := global.GS_DB.Model(&demo.Facility{})
	// json, err := json.MarshalIndent(info, "", " ")
	// global.GS_LOG.Info(string(json))

	// 如果有条件搜索 下方会自动创建搜索语句

	db := builder.Adapter("facility", global.GS_DB).
		Model("facility", &demo.Facility{}).
		Where("facility", "name", "like", info.Name).
		Where("facility", "code", "like", info.Code).
		Where("facility", "type", "in", info.Type).
		Where("facility", "status", "in", strings.Split(info.Status, ",")).
		Where("facility", "alarm_time", ">=", info.StartTime).
		Where("facility", "alarm_time", "<=", info.EndTime).
		Page("facility", info.Page, info.PageSize).
		Order("facility", info.Keyword).
		Go("facility")

	err = db.Count(&total).Find(&facilityList).Error
	if err != nil {
		global.GS_LOG.Error("GetFacilityInfoList failed", zap.Error(err))
	}

	return facilityList, total, err
}
