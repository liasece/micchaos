package boxes

import (
	"command"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/msg"
	"go.mongodb.org/mongo-driver/bson"
)

type Player struct {
	mod     *module.BaseModule
	Account `json:"account"`
	*log.Logger

	Session Session `json:"-"`
	Name    string  `json:"name"`
}

func (this *Player) Init(mod *module.BaseModule) {
	this.mod = mod
}

// 获取数据库主键
func (this *Player) GetPrimaryKey() bson.M {
	return bson.M{
		"account.uuid": this.Account.UUID,
	}
}

// after loaded from database
func (this *Player) AfterLoad() {
	this.Info("从数据库加载成功 %s", this.UUID)
}

func (this *Player) AfterOnline(session Session) {
	this.Info("登陆成功 %s", this.UUID)
	// Initial connect session
	this.Session = session

	send := &command.SC_ResEnterGame{}
	this.SendMsg(send)
}

func (this *Player) SendMsg(msgstr msg.MsgStruct) {
	if this.Session == nil {
		this.Debug("this.Session == nil")
		return
	}
	this.mod.SendMsgToClient(this.Session.Get("gate"),
		this.Session.Get("connectid"), msgstr)
}
