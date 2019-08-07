package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewClient(uri string) (*MongoClient, error) {
	res := &MongoClient{}
	err := res.Connect(uri)
	return res, err
}

type MongoClient struct {
	*mongo.Client
}

func (this *MongoClient) Connect(uri string) (err error) {
	opts := options.Client().ApplyURI(uri)
	this.Client, err = mongo.NewClient(opts)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = this.Client.Connect(ctx)
	return
}

func (this *MongoClient) Database(name string) *Database {
	res := &Database{}
	res.Database = this.Client.Database(name)
	return res
}
