package main

import (
	"gate"
	"github.com/liasece/micserver/conf"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/util"
	"os"
	"path/filepath"
	"sync"
)

type InitManager struct {
	modules    []module.IModule
	configPath string

	mutex sync.Mutex
}

var g_InitManager *InitManager
var g_InitManager_Lock sync.Once

func GetInitManger() *InitManager {
	g_InitManager_Lock.Do(func() {
		g_InitManager = &InitManager{}
	})
	return g_InitManager
}

func (this *InitManager) GetProgramModuleList() []module.IModule {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.modules == nil {
		this.modules = make([]module.IModule, 0)
		config := conf.TopConfig{}
		config.InitParse()
		isDevelopment := true
		for _, pid := range config.GetArgvModuleList() {
			isDevelopment = false
			stype := util.GetServerIDType(pid)
			log.Debug("App 初始化 ServerType[%s] ServerID[%s]", stype, pid)
			switch stype {
			case "gate":
				this.modules = append(this.modules, &gate.GatewayModule{BaseModule: module.BaseModule{ModuleID: pid}})
			}
		}
		if isDevelopment {
			this.modules = append(this.modules, &gate.GatewayModule{BaseModule: module.BaseModule{ModuleID: "gate001"}})
			this.modules = append(this.modules, &gate.GatewayModule{BaseModule: module.BaseModule{ModuleID: "gate002"}})
		}
	}
	return this.modules
}

func (this *InitManager) GetConfigPath() string {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.configPath == "" {
		pwd, err := os.Getwd()
		if err == nil {
			this.configPath = filepath.Join(pwd, "config", "config.json")
		} else {
			log.Error("os.Getwd() err:%v", err)
			return ""
		}
	}
	return this.configPath
}
