package file

import (
	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/request"
	"github.com/zhangrt/voyager1_platform/model/common/response"

	"github.com/zhangrt/voyager1_platform/model/file"
	file_res "github.com/zhangrt/voyager1_platform/model/file/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

// @Tags FileUploadAndDownload
// @Summary 上传文件示例
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {object} response.Response{data=file_res.FileResponse,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/uploadFile [post]
func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file file.FileUploadAndDownload
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GS_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileUploadAndDownloadService.UploadFile(c, header) // 文件上传后拿到文件路径
	if err != nil {
		global.GS_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(file_res.FileResponse{File: file}, "上传成功", c)
}

// EditFileName 编辑文件名或者备注
// @Tags FileUploadAndDownload
// @Summary 编辑文件名或者备注
// @accept multipart/form-data
// @Produce  application/json
// @Param data body file.FileUploadAndDownload true "编辑文件名或者备注"
// @Success 200 {object} response.Response{msg=string} "操作结果"
// @Router /fileUploadAndDownload/editFileName [post]
func (b *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file file.FileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := fileUploadAndDownloadService.EditFileName(file); err != nil {
		global.GS_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// @Tags FileUploadAndDownload
// @Summary 删除文件
// @Produce  application/json
// @Param data body file.FileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {object} response.Response{msg=string} "删除文件"
// @Router /fileUploadAndDownload/deleteFile [post]
func (b *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file file.FileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := fileUploadAndDownloadService.DeleteFile(c, file); err != nil {
		global.GS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// @Tags FileUploadAndDownload
// @Summary 分页文件列表
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router /fileUploadAndDownload/getFileList [post]
func (b *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
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
