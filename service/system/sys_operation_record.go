package system

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/system"
	systemReq "github.com/zhangrt/voyager1_platform/model/system/request"
)

//@author: [granty1](https://github.com/granty1)
//@function: CreateVo1OperationRecord
//@description: 创建记录
//@param: Vo1OperationRecord model.Vo1OperationRecord
//@return: err error

type OperationRecordService struct{}

func (operationRecordService *OperationRecordService) CreateVo1OperationRecord(Vo1OperationRecord system.Vo1OperationRecord) (err error) {
	err = global.GS_DB.Create(&Vo1OperationRecord).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteVo1OperationRecordByIds
//@description: 批量删除记录
//@param: ids request.IdsReq
//@return: err error

func (operationRecordService *OperationRecordService) DeleteVo1OperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.GS_DB.Delete(&[]system.Vo1OperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteVo1OperationRecord
//@description: 删除操作记录
//@param: Vo1OperationRecord model.Vo1OperationRecord
//@return: err error

func (operationRecordService *OperationRecordService) DeleteVo1OperationRecord(Vo1OperationRecord system.Vo1OperationRecord) (err error) {
	err = global.GS_DB.Delete(&Vo1OperationRecord).Error
	return err
}

//@author: [granty1](https://github.com/granty1)
//@function: DeleteVo1OperationRecord
//@description: 根据id获取单条操作记录
//@param: id uint
//@return: Vo1OperationRecord system.Vo1OperationRecord, err error

func (operationRecordService *OperationRecordService) GetVo1OperationRecord(id string) (Vo1OperationRecord system.Vo1OperationRecord, err error) {
	err = global.GS_DB.Where("id = ?", id).First(&Vo1OperationRecord).Error
	return
}

//@author: [granty1](https://github.com/granty1)
//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVo1OperationRecordInfoList
//@description: 分页获取操作记录列表
//@param: info systemReq.Vo1OperationRecordSearch
//@return: list interface{}, total int64, err error

func (operationRecordService *OperationRecordService) GetVo1OperationRecordInfoList(info systemReq.Vo1OperationRecordSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GS_DB.Model(&system.Vo1OperationRecord{})
	var Vo1OperationRecords []system.Vo1OperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&Vo1OperationRecords).Error
	return Vo1OperationRecords, total, err
}
