package manager

import (
	"ccmd"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/roc"
	"go.mongodb.org/mongo-driver/bson"
	"mongodb"
	"playermodule/boxes"
)

type PlayerDocManager struct {
	*log.Logger

	mod             *module.BaseModule
	mongo_userinfos *mongodb.UserInfos
	playerRoc       *roc.ROC
}

func (this *PlayerDocManager) Init(mod *module.BaseModule,
	userinfos *mongodb.UserInfos) {
	this.mod = mod
	this.mongo_userinfos = userinfos
	this.playerRoc = this.mod.GetROC(ccmd.ROCTypePlayer)
}

func (this *PlayerDocManager) getPlayerDoc(uuid string) *boxes.Player {
	if vi, ok := this.playerRoc.GetObj(uuid); ok {
		if p, ok := vi.(*boxes.Player); ok {
			return p
		}
	}
	return nil
}

func (this *PlayerDocManager) GetPlayerDoc(uuid string) *boxes.Player {
	return this.getPlayerDoc(uuid)
}

func (this *PlayerDocManager) loadOrStore(
	uuid string, p *boxes.Player) *boxes.Player {
	if vi, isLoad := this.playerRoc.GetOrRegObj(uuid, p); !isLoad {
		if p, ok := vi.(*boxes.Player); ok {
			this.Info("mod.ROCManager.RegObj OK")
			return p
		}
	}
	return p
}

// 从数据库获取用户信息
func (this *PlayerDocManager) getPlayerFromDB(uuid string) *boxes.Player {
	readPlayer := &boxes.Player{}
	err := this.mongo_userinfos.SelectOneByKey(bson.M{
		"account.uuid": uuid,
	}, readPlayer)
	if err != nil {
		this.Error("mongo_userinfos.SelectOneByKey err:%s", err.Error())
		return nil
	}
	readPlayer.Logger = this.Logger.Clone()
	readPlayer.Init(this.mod)
	readPlayer.AfterLoad()
	return readPlayer
}

// 必须取到用户数据，即使是从数据库取
func (this *PlayerDocManager) GetPlayerDocMust(uuid string) *boxes.Player {
	p := this.getPlayerDoc(uuid)
	if p == nil {
		p = this.getPlayerFromDB(uuid)
		p = this.loadOrStore(uuid, p)
	}
	return p
}

// 马上更新用户数据到数据库
func (this *PlayerDocManager) SavePlayerDocNow(player *boxes.Player) {
	_, err := this.mongo_userinfos.Update(player)
	if err != nil {
		this.Error("mongo_userinfos.Update err:%s", err.Error())
	}
}

// 向数据库中插入一个玩家
func (this *PlayerDocManager) InsertPlayerDocNow(player *boxes.Player) {
	_, err := this.mongo_userinfos.Upsert(player)
	if err != nil {
		this.Error("mongo_userinfos.Upsert err:%s", err.Error())
	}
}
