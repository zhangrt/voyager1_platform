package internal

import (
	"fmt"

	"github.com/zhangrt/voyager1_platform/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GS_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GS_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.GS_CONFIG.Pgsql.LogZap
	case "cockroach":
		logZap = global.GS_CONFIG.Cockroach.LogZap
	}
	if logZap {
		global.GS_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
