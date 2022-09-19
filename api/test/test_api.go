package test

import (
	"encoding/json"
	"os"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/common/response"
	"github.com/zhangrt/voyager1_platform/model/test"
	"github.com/zhangrt/voyager1_platform/utils"

	"github.com/zhangrt/voyager1_platform/model/test/request"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type TestApi struct{}

// @Tags TestApi
// @Summary 服务器信息
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取服务器信息"
// @Router /test/get [get]
func (e *TestApi) TestServer(c *gin.Context) {
	hostname, _ := os.Hostname()
	if server, err := systemConfigService.GetServerInfo(); err != nil {
		global.GS_LOG.Error(hostname, zap.Error(err))
		response.FailWithMessage(hostname, c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "Server:"+hostname, c)
	}
}

// @Tags TestApi
// @Summary 测试服务网格内部远程调用
// @Produce application/json
// @Param api query string true "通过api地址测试服务网格内部远程调用"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取服务器信息"
// @Router /test/server [get]
func (e *TestApi) TestMircoServer(c *gin.Context) {
	api := c.Query("api")
	hostname, _ := os.Hostname()
	prefix := utils.GetPrefixByServiceName("default", "facility-service")
	url := prefix + api
	client := resty.New()
	resp, _ := client.R().
		EnableTrace().
		Get(url)

	res := resp.Body()
	var data interface{}
	if err := json.Unmarshal(res, &data); err != nil {
		global.GS_LOG.Error("Unmarshal data faield", zap.Error(err))
	}

	response.OkWithDetailed(data, "Current Server:"+hostname, c)
}

// @Tags TestApi
// @Summary 测试 Get Api
// @Produce application/json
// @Param id query uint true "用id测试get请求"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /test/testGet [get]
func (e *TestApi) TestGet(c *gin.Context) {
	id := c.Query("id")

	if get_back, err := testService.TestGet(id); err != nil {
		global.GS_LOG.Error("get failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(get_back, c)
	}
}

// @Tags TestApi
// @Summary 测试 Post Api
// @accept application/json
// @Produce application/json
// @Param data body test.Test true "用id name测试post请求 redis"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /test/testPost [post]
func (e *TestApi) TestPost(c *gin.Context) {
	var info test.Test
	_ = c.ShouldBindJSON(&info)

	json, err := json.MarshalIndent(info, "", " ")
	global.GS_LOG.Info(string(json), zap.Error(err))

	if err := utils.Verify(info, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if get_back, err := testService.TestPost(info.ID, info.TestName); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(get_back, c)
	}
}

// @Tags TestApi
// @Summary 测试 Put Api
// @accept application/json
// @Produce application/json
// @Param data body test.Test true "测试put请求"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /test/testPut [put]
func (e *TestApi) TestPut(c *gin.Context) {
	var info test.Test
	_ = c.ShouldBindJSON(&info)

	if err := utils.Verify(info, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if get_back, err := testService.TestPost(info.ID, info.TestName); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(get_back, c)
	}
}

// @Tags TestApi
// @Summary 测试 Delete Api
// @accept application/json
// @Produce application/json
// @Param id query uint true "用id测试delete请求"
// @Success 200 {object} response.Response{msg=string} "okay"
// @Router /test/testDelete [delete]
func (e *TestApi) TestDelete(c *gin.Context) {
	id := c.Query("id")
	global.GS_LOG.Info("id=%d", zap.String("id", id))
	if get_back, err := testService.TestGet(id); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage(get_back, c)
	}
}

// @Tags TestApi
// @Summary 测试分页获取test列表
// @accept application/json
// @Produce application/json
// @Param data query request.TestSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取Test列表,返回包括列表,总数,页码,每页数量"
// @Router /test/getTestList [get]
func (s *TestApi) GetTestList(c *gin.Context) {
	var pageInfo request.TestSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testService.GetTestInfoList(pageInfo); err != nil {
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
