package boxes

import (
	"github.com/liasece/micserver/log"
	"go.mongodb.org/mongo-driver/bson"
)

type Player struct {
	Account `json:"account"`
	*log.Logger

	Session Session `json:"-"`
	Name    string  `json:"name"`
}

func (this *Player) GetPrimaryKey() bson.M {
	return bson.M{
		"account.uuid": this.Account.UUID,
	}
}

func (this *Player) AfterLoad() {
	this.Info("加载成功 %s", this.UUID)
}

func (this *Player) AfterOnline(session Session) {
	this.Info("登陆成功 %s", this.UUID)
	// 初始化会话
	this.Session = session
}
