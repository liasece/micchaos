package playermodule

import (
	"fmt"
	"github.com/liasece/micserver/module"
	"mongodb"
	// "playermodule/boxes"
	"playermodule/manager"
)

type PlayerModule struct {
	module.BaseModule

	PlayerDocManager manager.PlayerDocManager
	mongo_userinfos  *mongodb.UserInfos
	HandlerClient    HandlerClient
	HandlerServer    HandlerServer
}

func (this *PlayerModule) AfterInitModule() {
	this.BaseModule.AfterInitModule()

	this.HandlerClient.Init(this)
	this.HandlerServer.Init(this)

	// 数据库初始化
	mongouri := this.Configer.GetSetting("mongodb")
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

	// player := &boxes.Player{
	// 	Account: boxes.Account{
	// 		UUID: "13412341",
	// 	},
	// 	Name: "jansen",
	// }
	// _, err := this.mongo_userinfos.Upsert(player)
	// if err != nil {
	// 	this.Error("mongo_userinfos.Upsert err:%s", err.Error())
	// }

	this.PlayerDocManager.Init(&this.BaseModule, this.mongo_userinfos)
	this.PlayerDocManager.Logger = this.Logger

	this.RegForwardToServer(this.HandlerServer.OnRecvServerMsg)
	this.RegForwardFromGate(this.HandlerServer.OnRecvGateMsg)
}
