package boxes

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Player struct {
	session map[string]string

	Account `json:"account"`
	Name    string `json:"name"`
}

func (this *Player) GetPrimaryKey() bson.M {
	return bson.M{
		"account.uuid": this.Account.UUID,
	}
}

func (this *Player) GetSession() map[string]string {
	return this.session
}

func (this *Player) MergeSession(session map[string]string) {
	if this.session == nil {
		this.session = make(map[string]string)
	}
	for k, v := range session {
		this.session[k] = v
	}
}
