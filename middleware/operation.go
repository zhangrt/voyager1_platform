package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/zhangrt/voyager1_platform/global"
	"github.com/zhangrt/voyager1_platform/model/system"
	"github.com/zhangrt/voyager1_platform/service"

	auth "github.com/zhangrt/voyager1_core/auth/luna"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId string
		body = ReadAll(c, body)
		userId = GetUserId(c, userId)
		record := system.Vo1OperationRecord{
			Ip:       c.ClientIP(),
			Method:   c.Request.Method,
			Path:     c.Request.URL.Path,
			Agent:    c.Request.UserAgent(),
			Body:     string(body),
			PersonId: userId,
		}

		// 上传文件时候 中间件日志进行裁断操作
		if strings.Index(c.GetHeader("Content-Type"), "multipart/form-data") > -1 {
			if len(record.Body) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Body)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if strings.Index(c.Writer.Header().Get("Pragma"), "public") > -1 ||
			strings.Index(c.Writer.Header().Get("Expires"), "0") > -1 ||
			strings.Index(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/force-download") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/octet-stream") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/download") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Disposition"), "attachment") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") > -1 {
			if len(record.Resp) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Resp)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		if err := operationRecordService.CreateVo1OperationRecord(record); err != nil {
			global.GS_LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func ReadAll(c *gin.Context, body []byte) []byte {
	if c.Request.Method != http.MethodGet {
		var err error
		body, err = ioutil.ReadAll(c.Request.Body)
		if err != nil {
			global.GS_LOG.Error("read body from request error:", zap.Error(err))
		} else {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	} else {
		query := c.Request.URL.RawQuery
		query, _ = url.QueryUnescape(query)
		split := strings.Split(query, "&")
		m := make(map[string]string)
		for _, v := range split {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
		body, _ = json.Marshal(&m)
	}
	return body
}

func GetUserId(c *gin.Context, userId string) string {
	claims, _ := auth.GetClaims(c)
	if claims.ID != "" {
		userId = claims.ID
	} else {
		userId = c.Request.Header.Get(global.GS_CONFIG.AUTHKey.UserId)
	}
	return userId
}
