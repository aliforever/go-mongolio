package mongolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) UpdateByID(id any, data interface{}, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateByID(context.Background(), id, bson.M{"$set": data}, options...)
}

func (c *C[T]) UpdateCustom(filter bson.M, data interface{}, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
}

func (c *C[T]) UpdateCustomMany(filter bson.M, data interface{}, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateMany(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
}
