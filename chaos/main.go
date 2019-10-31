package main

import (
	"github.com/liasece/micserver"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/util/monitor"
)

func main() {
	// 初始化性能监控
	monitor.BindPprof("", 8888)

	// 初始化 MicServer
	app, err := micserver.CreateApp(GetInitManger().GetConfigPath(),
		GetInitManger().GetProgramModuleList())
	if err != nil {
		log.Fatal("Create app fatal: %v", err)
		return
	}

	// app 开始运行 阻塞
	app.RunAndBlock()
}
