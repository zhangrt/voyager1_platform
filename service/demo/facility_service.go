package demo

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/demo"
	"github.com/zhangrt/voyager1_platform/model/demo/request"
	"github.com/zhangrt/voyager1_platform/model/demo/response"

	"gorm.io/gorm"
)

type FacilityService struct{}

func (facilityService *FacilityService) AddFacility(f demo.Facility) (data demo.Facility, err error) {
	var facility demo.Facility
	if global.GS_DB == nil {
		return f, err
	}
	if !errors.Is(global.GS_DB.Where("code = ?", f.Code).First(&facility).Error, gorm.ErrRecordNotFound) { // 判断设备唯一标识是否已存在
		return data, errors.New("该设备已存在")
	}
	err = global.GS_DB.Create(&f).Error
	return f, err
}

func (facilityService *FacilityService) RemoveFacility(f demo.Facility) (data demo.Facility, err error) {
	if global.GS_DB == nil {
		return f, err
	}
	err = global.GS_DB.Delete(&f).Error // 模型有DeletedAt字段，将自动软删除
	// db.Unscoped().Delete(&f) // 永久删除
	return f, err
}

func (facilityService *FacilityService) UpdateFacility(f demo.Facility) (data demo.Facility, err error) {
	err = global.GS_DB.Model(&f).Where("deleted_at is Null").Updates(&f).Error
	return f, err
}

func (facilityService *FacilityService) GetFacility(id string) (data demo.Facility, err error) {
	facility, err := facilityRepository.GetFacility(context.Background(), id)
	return facility, err
}

func (facilityService *FacilityService) GetFacilityInfoList(info request.FacilitySearch) (list interface{}, total int64, err error) {
	var facilityList []response.FacilityResponse = []response.FacilityResponse{}
	var data []demo.Facility
	data, total, err = facilityRepository.GetFacilityInfoList(info)
	var lens = 6 // 6种设备类型
	var types []string
	if info.Type != "" {
		types = strings.Split(info.Type, ",")
		lens = len(types)
	}

	// 按不同类型分组返回
	for t := 0; t < lens; t++ {
		var Type string
		if lens == 6 {
			Type = strconv.Itoa(t)
		} else {
			Type = types[t]
		}
		var list []demo.Facility = []demo.Facility{}
		for index := range data {
			if Type == data[index].Type {
				list = append(list, data[index])
			}
		}
		if len(list) > 0 {
			facilityList = append(facilityList, response.FacilityResponse{Type: Type, FacilityList: list})
		}
	}
	return facilityList, total, err
}
