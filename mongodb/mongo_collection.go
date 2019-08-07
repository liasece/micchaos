package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var optUpset = options.Update().SetUpsert(true)

type Collection struct {
	*mongo.Collection
}

func (this *Collection) InsertOne(
	doc interface{}) (*mongo.InsertOneResult, error) {
	return this.Collection.InsertOne(context.Background(), doc)
}

func (this *Collection) InsertMany(
	documents []interface{}) (*mongo.InsertManyResult, error) {
	return this.Collection.InsertMany(context.Background(), documents)
}

func (this *Collection) UpdateOrInsertOne(
	filter interface{}, doc interface{}) (*mongo.UpdateResult, error) {
	return this.Collection.UpdateOne(context.Background(), filter, doc,
		options.Update().SetUpsert(true))
}

func (this *Collection) UpdateOne(
	filter interface{}, doc interface{}) (*mongo.UpdateResult, error) {
	return this.Collection.UpdateOne(context.Background(), filter, doc)
}

func (this *Collection) UpdateMany(
	filter interface{}, doc interface{}) (*mongo.UpdateResult, error) {
	return this.Collection.UpdateOne(context.Background(), filter, doc)
}

func (this *Collection) SelectOne(
	filter interface{}) *mongo.SingleResult {
	return this.Collection.FindOne(context.Background(), filter)
}
