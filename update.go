package mongolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) UpdateByID(id any, data interface{}, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateByID(context.Background(), id, bson.M{"$set": data}, options...)
	return
}

func (c *C[T]) UpdateCustom(filter bson.M, data interface{}, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateOne(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
	return
}

func (c *C[T]) UpdateCustomMany(filter bson.M, data interface{}, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateMany(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
	return
}
