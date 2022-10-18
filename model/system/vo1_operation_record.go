// 自动生成模板SysOperationRecord
package system

import (
	"time"

	"github.com/zhangrt/voyager1_core/global"
)

// 如果含有time.Time 请自行import time包
type Vo1OperationRecord struct {
	global.GS_BASE_MODEL_ID_STRING
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                   // 请求ip
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                       // 请求方法
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                       // 请求状态
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                            // 代理
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`  // 错误信息
	Body         string        `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"`                 // 请求Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"`                 // 响应Body
	PersonId     string        `json:"person_id" form:"user_id" gorm:"column:user_id;comment:用户id"`                  // 用户id
	Person       Vo1Person     `json:"person"`
}

func (Vo1OperationRecord) TableName() string {
	return "vo1_operation_record"
}
