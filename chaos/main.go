package main

import (
	"github.com/liasece/micserver"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/util"
)

func main() {
	util.BindPprof("", 8888)
	app, err := micserver.CreateApp(GetInitManger().GetConfigPath(),
		GetInitManger().GetProgramModuleList())
	if err != nil {
		log.Fatal("Create app fatal: %v", err)
		return
	}

	app.Run()
}
