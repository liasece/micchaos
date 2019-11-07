package playermodule

import (
	"fmt"
	"mongodb"

	"ccmd"
	"github.com/liasece/micserver/module"
	"playermodule/manager"
)

type PlayerModule struct {
	module.BaseModule

	PlayerDocManager manager.PlayerDocManager
	mongo_userinfos  *mongodb.UserInfos
	HandlerClient    HandlerClient
	HandlerServer    HandlerServer
}

func NewPlayerModule(moduleid string) *PlayerModule {
	res := &PlayerModule{}
	res.BaseModule.SetModuleID(moduleid)
	return res
}

func (this *PlayerModule) AfterInitModule() {
	this.BaseModule.AfterInitModule()

	// 初始化业务逻辑
	this.BaseModule.ROCManager.NewObjectType("Player")
	// 事件处理器
	this.HandlerClient.Init(this)
	this.HandlerServer.Init(this)

	// 数据库初始化
	mongouri := this.Configer.GetString(ccmd.ConfMongoDB)
	if mongouri != "" {
		this.Debug("连接 MondgoDB[%s]", mongouri)
		var err error

		// 初始化玩家数据表
		this.mongo_userinfos, err = mongodb.NewUserInfos(this, mongouri)
		if err != nil {
			this.Error("mongodb.NewUserInfos err: %s", err.Error())
			panic(fmt.Sprintf("mongodb.NewUserInfos err: %s", err.Error()))
		} else {
			this.Debug("mongodb.NewUserInfos scesse")
		}
	}

	this.PlayerDocManager.Init(&this.BaseModule, this.mongo_userinfos)
	this.PlayerDocManager.Logger = this.Logger

	this.HookServer(&this.HandlerServer)
}
