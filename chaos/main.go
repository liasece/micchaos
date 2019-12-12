package main

import (
	"time"

	"github.com/liasece/micserver"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/util/monitor"
)

func main() {
	// 初始化 MicServer
	app, err := micserver.SetupApp(GetInitManger().GetConfigPath())
	if err != nil {
		log.Fatal("Create app fatal: %v", err)
		time.Sleep(time.Second * 1)
		return
	}

	log.Info("即将运行")
	// 初始化性能监控
	monitor.BindPprof("", 8888)

	// app 开始运行 阻塞
	app.RunAndBlock(GetInitManger().GetProgramModuleList())
}
