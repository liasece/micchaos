package boxes

import (
	"command"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/session"
	"go.mongodb.org/mongo-driver/bson"
)

type Player struct {
	mod     *module.BaseModule
	Account `json:"account"`
	*log.Logger

	Session session.Session `json:"-"`
	Name    string          `json:"name"`
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

func (this *Player) AfterOnline(session session.Session) {
	this.Info("登陆成功 %s", this.UUID)
	// Initial connect session
	this.Session = session

	send := &command.SC_ResEnterGame{}
	this.SendMsg(send)
}

func (this *Player) SendMsg(msg interface{}) {
	btop := command.GetSCTopLayer(msg)
	this.mod.SendBytesToClient(this.Session.GetBindServer("gate"),
		this.Session.GetConnectID(), 0, btop)
}
