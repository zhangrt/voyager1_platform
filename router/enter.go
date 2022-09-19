package router

import (
	"github.com/zhangrt/voyager1_platform/router/demo"
	file "github.com/zhangrt/voyager1_platform/router/file"
	"github.com/zhangrt/voyager1_platform/router/statistics"
	system "github.com/zhangrt/voyager1_platform/router/system"
	test "github.com/zhangrt/voyager1_platform/router/test"
)

type RouterGroup struct {
	Test       test.RouterGroup
	File       file.RouterGroup
	System     system.RouterGroup
	Demo       demo.RouterGroup
	Statistics statistics.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
