package main

import (
	"gatemodule"
	"github.com/liasece/micserver/conf"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/util"
	"os"
	"path/filepath"
	"sync"
)

type InitManager struct {
	modules    map[string]module.IModule
	configPath string

	hasInit bool
	mutex   sync.Mutex
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

	if !this.hasInit {
		this.hasInit = true
		this.modules = make(map[string]module.IModule)
		config := conf.TopConfig{}
		config.InitParse()
		isDevelopment := true

		// 遍历所有的参数指定的模块名
		for _, pid := range config.GetArgvModuleList() {
			isDevelopment = false
			stype := util.GetServerIDType(pid)
			log.Debug("App 初始化 ServerType[%s] ServerID[%s]", stype, pid)
			switch stype {
			case "gate":
				this.addModule(&gatemodule.GatewayModule{
					BaseModule: module.BaseModule{
						ModuleID: pid,
					},
				})
			}
		}

		// 如果当前是开发模式，添加如下的列表
		if isDevelopment {
			this.addModule(&gatemodule.GatewayModule{
				BaseModule: module.BaseModule{
					ModuleID: "gate001",
				},
			})
			this.addModule(&gatemodule.GatewayModule{
				BaseModule: module.BaseModule{
					ModuleID: "gate002",
				},
			})
		}
	}
	return this.getModuleSlice()
}

// 添加一个模块
func (this *InitManager) addModule(module module.IModule) bool {
	if this.modules == nil {
		return false
	}
	this.modules[module.GetModuleID()] = module
	return true
}

// 获取模块列表的切片形式
func (this *InitManager) getModuleSlice() []module.IModule {
	res := make([]module.IModule, 0)
	for _, m := range this.modules {
		res = append(res, m)
	}
	return res
}

// 获取配置文件的路径
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
