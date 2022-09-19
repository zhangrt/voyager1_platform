package file

import (
	v1 "github.com/zhangrt/voyager1_platform/api"

	"github.com/gin-gonic/gin"
)

type FileRouter struct{}

func (e *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("fileUploadAndDownload")
	exaTestApi := v1.ApiGroupApp.FileApiGroup.FileUploadAndDownloadApi
	{
		testRouter.POST("uploadFile", exaTestApi.UploadFile)
		testRouter.POST("deleteFile", exaTestApi.DeleteFile)
		testRouter.POST("getFileList", exaTestApi.GetFileList)
		testRouter.POST("editFileName", exaTestApi.EditFileName)
	}
}
