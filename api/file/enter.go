package file

import "github.com/zhangrt/voyager1_platform/service"

type ApiGroup struct {
	FileUploadAndDownloadApi
}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.FileServiceGroup.FileUploadAndDownloadService
)
