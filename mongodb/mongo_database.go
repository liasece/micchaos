package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	*mongo.Database
}

func (this *Database) Collection(name string) *Collection {
	res := &Collection{}
	res.Collection = this.Database.Collection(name)
	return res
}
