package playermodule

import (
	"ccmd"
	"fmt"
	"mongodb"
	"playermodule/manager"

	"github.com/liasece/micserver/module"
)

type PlayerModule struct {
	module.BaseModule

	PlayerDocManager manager.PlayerDocManager
	mongoUserinfos   *mongodb.UserInfos
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
	this.BaseModule.NewROC(ccmd.ROCTypePlayer)
	// 事件处理器
	this.HandlerClient.Init(this)
	this.HandlerServer.Init(this)

	// 数据库初始化
	mongouri := this.GetConfiger().GetString(ccmd.ConfMongoDB)
	if mongouri != "" {
		this.Debug("连接 MondgoDB[%s]", mongouri)
		var err error

		// 初始化玩家数据表
		this.mongoUserinfos, err = mongodb.NewUserInfos(this, mongouri)
		if err != nil {
			this.Error("mongodb.NewUserInfos err: %s", err.Error())
			panic(fmt.Sprintf("mongodb.NewUserInfos err: %s", err.Error()))
		} else {
			this.Debug("mongodb.NewUserInfos scesse")
		}
	}

	this.PlayerDocManager.Init(&this.BaseModule, this.mongoUserinfos)
	this.PlayerDocManager.Logger = this.Logger

	this.HookServer(&this.HandlerServer)
}
