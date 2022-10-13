package system

import (
	"errors"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/model/system/request"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteVo1Dictionary
//@description: 创建字典数据
//@param: Vo1Dictionary model.Vo1Dictionary
//@return: err error

type DictionaryService struct{}

func (dictionaryService *DictionaryService) CreateVo1Dictionary(Vo1Dictionary system.Vo1Dictionary) (err error) {
	if (!errors.Is(global.GS_DB.First(&system.Vo1Dictionary{}, "id = ?", Vo1Dictionary.ID).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的ID(key)，不允许创建")
	}
	err = global.GS_DB.Create(&Vo1Dictionary).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteVo1Dictionary
//@description: 删除字典数据
//@param: Vo1Dictionary model.Vo1Dictionary
//@return: err error

func (dictionaryService *DictionaryService) DeleteVo1Dictionary(Vo1Dictionary system.Vo1Dictionary) (err error) {
	err = global.GS_DB.Where("id = ?", Vo1Dictionary.ID).Preload("Vo1DictionaryDetails").First(&Vo1Dictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = global.GS_DB.Delete(&Vo1Dictionary).Error
	if err != nil {
		return err
	}

	if Vo1Dictionary.Dictionarys != nil {
		return global.GS_DB.Where("parent_id=?", Vo1Dictionary.ID).Delete(Vo1Dictionary.Dictionarys).Error
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateVo1Dictionary
//@description: 更新字典数据
//@param: Vo1Dictionary *model.Vo1Dictionary
//@return: err error

func (dictionaryService *DictionaryService) UpdateVo1Dictionary(Vo1Dictionary *system.Vo1Dictionary) (err error) {
	var dict system.Vo1Dictionary
	Vo1DictionaryMap := map[string]interface{}{
		"NameCN":      Vo1Dictionary.NameCN,
		"NameEN":      Vo1Dictionary.NameEN,
		"Key":         Vo1Dictionary.ID,
		"Description": Vo1Dictionary.Description,
	}
	db := global.GS_DB.Where("id = ?", Vo1Dictionary.ID).First(&dict)
	if dict.ID != Vo1Dictionary.ID {
		if !errors.Is(global.GS_DB.First(&system.Vo1Dictionary{}, "id = ?", Vo1Dictionary.ID).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的ID(Key)，不允许创建")
		}
	}
	err = db.Updates(Vo1DictionaryMap).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVo1Dictionary
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, Vo1Dictionary model.Vo1Dictionary

func (dictionaryService *DictionaryService) GetVo1Dictionary(NameEN string, Id string) (Vo1Dictionary system.Vo1Dictionary, err error) {
	err = global.GS_DB.Where("id = ? OR NameEN = ?", Id, NameEN).First(&Vo1Dictionary).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetVo1DictionaryInfoList
//@description: 分页获取字典列表
//@param: info request.Vo1DictionarySearch
//@return: err error, list interface{}, total int64

func (dictionaryService *DictionaryService) GetVo1DictionaryInfoList(info request.Vo1DictionarySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GS_DB.Model(&system.Vo1Dictionary{})
	var Vo1Dictionarys []system.Vo1Dictionary
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.NameCN != "" {
		db = db.Where("name LIKE ?", "%"+info.NameCN+"%")
	}
	if info.ID != "" {
		db = db.Where("type LIKE ?", "%"+info.ID+"%")
	}
	if info.NameEN != "" {
		db = db.Where("status LIKE ?", "%"+info.NameEN+"%")
	}
	if info.Description != "" {
		db = db.Where("description LIKE ?", "%"+info.Description+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&Vo1Dictionarys).Error
	return Vo1Dictionarys, total, err
}
