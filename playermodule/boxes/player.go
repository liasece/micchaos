package boxes

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Player struct {
	Account `json:"account"`
	Name    string `json:"name"`
}

func (this *Player) GetPrimaryKey() bson.M {
	return bson.M{
		"account.uuid": this.Account.UUID,
	}
}
