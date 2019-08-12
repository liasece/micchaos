package mongodb

import (
	"fmt"
	"github.com/liasece/micserver/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUser interface {
	GetPrimaryKey() bson.M
}

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

func (this *UserInfos) Upsert(obj IUser) (*mongo.UpdateResult, error) {
	bsonm, err := GetBsonProxyJsonByObj(obj)
	if err != nil {
		return nil, err
	}
	return this.UpdateOrInsertOne(obj.GetPrimaryKey(),
		bson.M{
			"$set": bsonm,
		},
	)
}

func (this *UserInfos) Update(obj IUser) (*mongo.UpdateResult, error) {
	bsonm, err := GetBsonProxyJsonByObj(obj)
	if err != nil {
		return nil, err
	}
	return this.UpdateOne(obj.GetPrimaryKey(),
		bson.M{
			"$set": bsonm,
		},
	)
}

func (this *UserInfos) SelectOne(obj IUser) error {
	res := this.Collection.SelectOne(obj.GetPrimaryKey())
	var resBson = bson.M{}
	err := res.Decode(&resBson)
	if err != nil {
		return err
	}
	return GetObjProxyJsonByBson(resBson, obj)
}

func (this *UserInfos) SelectOneByKey(primarykey bson.M, obj interface{}) error {
	res := this.Collection.SelectOne(primarykey)
	var resBson = bson.M{}
	err := res.Decode(&resBson)
	if err != nil {
		return err
	}
	return GetObjProxyJsonByBson(resBson, obj)
}

func (this *UserInfos) SelectOneByAccount(account string,
	obj interface{}) error {
	primarykey := bson.M{
		"$or": bson.A{
			bson.M{"account.phonenumber": account},
			bson.M{"account.uuid": account},
			bson.M{"account.loginname": account},
		},
	}
	res := this.Collection.SelectOne(primarykey)
	var resBson = bson.M{}
	err := res.Decode(&resBson)
	if err != nil {
		return err
	}
	return GetObjProxyJsonByBson(resBson, obj)
}

func (this *UserInfos) FindOneOrCreate(account string, newobj interface{},
	result interface{}) error {
	primarykey := bson.M{
		"$or": bson.A{
			bson.M{"account.phonenumber": account},
			bson.M{"account.uuid": account},
			bson.M{"account.loginname": account},
		},
	}
	bsonm, err := GetBsonProxyJsonByObj(newobj)
	if err != nil {
		return err
	}
	res := this.Collection.FindOneOrCreate(primarykey, bsonm)
	var resBson = bson.M{}
	err = res.Decode(&resBson)
	if err != nil {
		return err
	}
	return GetObjProxyJsonByBson(resBson, result)
}
