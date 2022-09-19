package response

import (
	"github.com/zhangrt/voyager1_platform/model/file"
)

type FileResponse struct {
	File file.FileUploadAndDownload `json:"file"`
}
