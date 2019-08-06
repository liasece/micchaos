package main

import (
	"gate"
	"github.com/liasece/micserver"
	"github.com/liasece/micserver/conf"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/util"
	"os"
	"path/filepath"
)

func getProgramModuleList() []module.IModule {
	res := make([]module.IModule, 0)
	config := conf.TopConfig{}
	config.InitParse()
	isDevelopment := true
	for _, pid := range config.GetArgvModuleList() {
		isDevelopment = false
		stype := util.GetServerIDType(pid)
		log.Debug("App 初始化 ServerType[%s] ServerID[%s]", stype, pid)
		switch stype {
		case "gate":
			res = append(res, &gate.GatewayModule{BaseModule: module.BaseModule{ModuleID: pid}})
		}
	}
	if isDevelopment {
		res = append(res, &gate.GatewayModule{BaseModule: module.BaseModule{ModuleID: "gate001"}})
		res = append(res, &gate.GatewayModule{BaseModule: module.BaseModule{ModuleID: "gate002"}})
	}

	return res
}

func main() {
	configpath := ""
	pwd, err := os.Getwd()
	if err == nil {
		configpath = filepath.Join(pwd, "config", "config.json")
	} else {
		log.Error("os.Getwd() err:%v", err)
		return
	}
	util.BindPprof("", 8888)
	app, err := micserver.CreateApp(configpath, getProgramModuleList())
	if err != nil {
		log.Fatal("Create app fatal: %v", err)
		return
	}

	app.Run()
}
