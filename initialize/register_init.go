package initialize

import (
	_ "github.com/zhangrt/voyager1_platform/source/system"
	_ "github.com/zhangrt/voyager1_platform/source/test"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
