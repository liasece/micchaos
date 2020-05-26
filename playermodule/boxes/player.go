package boxes

import (
	"github.com/liasece/micchaos/ccmd"
	"github.com/liasece/micserver/log"
	"github.com/liasece/micserver/module"
	"github.com/liasece/micserver/roc"
	"github.com/liasece/micserver/session"
	"go.mongodb.org/mongo-driver/bson"
)

type Player struct {
	mod     *module.BaseModule
	Account `json:"account"`
	*log.Logger

	Session *session.Session `json:"-"`
	Name    string           `json:"name"`
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

func (this *Player) GetROCObjID() string {
	return this.Account.UUID
}

func (this *Player) GetROCObjType() roc.ObjType {
	return ccmd.ROCTypePlayer
}

func (this *Player) OnROCCall(path *roc.Path, arg []byte) ([]byte, error) {
	this.Info("ROC调用执行:[%s],%+v", path.String(), arg)
	switch path.Move() {
	case "GateClose":
		this.OnGateClose()
	}
	return nil, nil
}

func (this *Player) OnGateClose() {
	this.Info("Player连接关闭")
}

// after loaded from database
func (this *Player) AfterLoad() {
	this.Info("从数据库加载成功 %s", this.UUID)
	this.Logger.SetTopic("Player[" + this.UUID + "]")
}

func (this *Player) AfterOnline(session *session.Session) {
	this.Info("登陆成功 %s", this.UUID)
	// Initial connect session
	this.Session = session

	send := &ccmd.SC_ResEnterGame{}
	this.SendMsg(send)

	this.mod.ROCCallNR(roc.O(ccmd.ROCTypePlayer, this.UUID).F("Regdata"),
		[]byte{12, 43, 78, 222, 96})
}

func (this *Player) SendMsg(msg interface{}) {
	btop := ccmd.GetSCTopLayer(msg)
	this.Session.SendMsg(this.mod, "gate", 0, btop)
}
