package loginmodule

import (
	"fmt"
	"github.com/liasece/micserver/module"
	"mongodb"
)

type LoginModule struct {
	module.BaseModule

	mongo_userinfos *mongodb.UserInfos
	HandlerClient   HandlerClient
	HandlerServer   HandlerServer
}

func (this *LoginModule) AfterInitModule() {
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

	// 系统事件监听初始化
	subnetManager := this.GetSubnetManager()
	if subnetManager != nil {
		subnetManager.RegHandleServerMsg(this.HandlerServer.OnRecvServerMsg)
		subnetManager.RegHandleGateMsg(this.HandlerServer.OnRecvGateMsg)
	}
}
