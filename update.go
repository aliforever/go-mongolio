package mongorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) Update(options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.Name()).UpdateByID(context.Background(), i.ID(), c, options...)
	return
}

func (c *C[T]) UpdateCustom(filter, data bson.M, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.Name()).UpdateOne(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
	return
}

func (c *C[T]) UpdateCustomMany(filter, data bson.M, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.Name()).UpdateMany(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
	return
}
