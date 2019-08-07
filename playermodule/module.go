package playermodule

import (
	"encoding/json"
	"fmt"
	"github.com/liasece/micserver/module"
	"go.mongodb.org/mongo-driver/bson"
	"mongodb"
	"playermodule/boxes"
)

type PlayerModule struct {
	module.BaseModule

	mongo_userinfos *mongodb.UserInfos
}

func (this *PlayerModule) AfterInitModule() {
	this.BaseModule.AfterInitModule()

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

	subnet := this.GetSubnetManager()
	if subnet != nil {

	}

	player := &boxes.Player{
		Account: boxes.Account{
			UUID: "13412341",
		},
		Name: "jansen",
	}
	_, err := this.mongo_userinfos.Upsert(player)
	if err != nil {
		this.Error("mongo_userinfos.Upsert err:%s", err.Error())
	}

	readPlayer := &boxes.Player{}
	err = this.mongo_userinfos.SelectOneByKey(bson.M{
		"account.uuid": "13412341",
	}, readPlayer)
	if err != nil {
		this.Error("mongo_userinfos.SelectOneByKey err:%s", err.Error())
	}
	jsonb, _ := json.Marshal(readPlayer)
	this.Info("%s", string(jsonb))
}
