package system

import (
	"github.com/zhangrt/voyager1_platform/config"
)

// 配置文件结构体
type Config struct {
	Config config.Server `json:"config"`
}
