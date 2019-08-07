package mongodb

import (
	"fmt"
	"github.com/liasece/micserver/module"
)

type UserInfos struct {
	*Collection

	database *Database
	client   *MongoClient
	mod      module.IModule
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
	if databasename == "" {
		return fmt.Errorf("empty database name")
	} else {
		this.database = this.client.Database(databasename)
	}
	collectionname := this.mod.GetConfiger().GetSetting("userinfos_collection")
	if collectionname == "" {
		return fmt.Errorf("empty collection name")
	} else {
		this.Collection = this.database.Collection(collectionname)
	}
	return nil
}
