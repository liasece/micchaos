package mongodb

import (
	"context"
	"fmt"
	"github.com/liasece/micserver/module"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var optUpset = options.Update().SetUpsert(true)

type UserInfos struct {
	*mongo.Collection

	client *MongoClient
	mod    module.IModule
}

func NewUserInfos(mod module.IModule, uri string) (*UserInfos, error) {
	client, err := NewClient(uri)
	if err != nil {
		return nil, err
	}
	userinfos := &UserInfos{}
	err1 := userinfos.Init(mod, client)
	return userinfos, err1
}

func (this *UserInfos) Init(mod module.IModule, client *MongoClient) error {
	this.client = client
	this.mod = mod

	databasename := this.mod.GetConfiger().GetSetting("database")
	collectionname := this.mod.GetConfiger().GetSetting("userdoc_collection")
	if databasename == "" || collectionname == "" {
		return fmt.Errorf("empty database name or collection name")
	} else {
		this.Collection = this.client.Database(databasename).
			Collection(collectionname)
	}
	return nil
}

func (this *UserInfos) InsertOne(
	doc interface{}) (*mongo.InsertOneResult, error) {
	return this.Collection.InsertOne(context.Background(), doc)
}

func (this *UserInfos) UpdateOrInsertOne(
	where interface{}, doc interface{}) (*mongo.UpdateResult, error) {
	return this.Collection.UpdateOne(context.Background(), where, doc,
		options.Update().SetUpsert(true))
}
