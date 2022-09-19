package file

import (
	"errors"
	"mime/multipart"
	"strings"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	model "github.com/zhangrt/voyager1_platform/model/file"

	upload "github.com/zhangrt/voyager1_core/oss"

	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadService struct{}

//@function: Upload
//@description: 创建文件上传记录
//@param: file model.FileUploadAndDownload
//@return: error

func (fileUploadAndDownloadService *FileUploadAndDownloadService) Upload(file model.FileUploadAndDownload) error {
	return global.GS_DB.Create(&file).Error
}

//@function: FindFile
//@description: 查询文件记录
//@param: key uint
//@return: model.FileUploadAndDownload, error

func (fileUploadAndDownloadService *FileUploadAndDownloadService) FindFile(id uint) (model.FileUploadAndDownload, error) {
	var file model.FileUploadAndDownload
	err := global.GS_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.FileUploadAndDownload
//@return: err error

func (fileUploadAndDownloadService *FileUploadAndDownloadService) DeleteFile(ctx *gin.Context, file model.FileUploadAndDownload) (err error) {
	var fileFromDb model.FileUploadAndDownload
	fileFromDb, err = fileUploadAndDownloadService.FindFile(file.ID)
	if err != nil {
		return
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(ctx, fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GS_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// EditFileName 编辑文件名或者备注
func (fileUploadAndDownloadService *FileUploadAndDownloadService) EditFileName(file model.FileUploadAndDownload) (err error) {
	var fileFromDb model.FileUploadAndDownload
	return global.GS_DB.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (fileUploadAndDownloadService *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.GS_DB.Model(&model.FileUploadAndDownload{})
	var fileLists []model.FileUploadAndDownload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

//@function: UploadFile
//@description: 根据配置文件判断是文件上传到Minio或其他OSS
//@param: ctx &gin.Context, header *multipart.FileHeader
//@return: file model.FileUploadAndDownload, err error

func (fileUploadAndDownloadService *FileUploadAndDownloadService) UploadFile(ctx *gin.Context, header *multipart.FileHeader) (file model.FileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(ctx, header)
	if uploadErr != nil {
		panic(err)
	}
	s := strings.Split(header.Filename, ".")
	f := model.FileUploadAndDownload{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	return f, fileUploadAndDownloadService.Upload(f)

}

func (fileUploadAndDownloadService *FileUploadAndDownloadService) DownloadFile(ctx *gin.Context, key string) error {
	oss := upload.NewOss()

	if oss != nil {
		oss.DownloadFile(ctx, key)
	}

	return nil
}
