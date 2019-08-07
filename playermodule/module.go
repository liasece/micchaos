package playermodule

import (
	"fmt"
	"github.com/liasece/micserver/module"
	"mongodb"
)

type PlayerModule struct {
	module.BaseModule

	mongo_userinfos *mongodb.UserInfos
}

func (this *PlayerModule) AfterInitModule() {
	this.BaseModule.AfterInitModule()
	this.Debug("PlayerModule init finish %s", this.GetModuleID())

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
}
