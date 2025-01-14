package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) UpdateByID(
	id any,
	data interface{},
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateByID(context.Background(), id, bson.M{"$set": data}, opts...)
}

func (c *C[T]) UpdateCustom(
	filter bson.M,
	data interface{},
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.M{
		"$set": data,
	}, opts...)
}

func (c *C[T]) UpdateCustomMany(
	filter bson.M,
	data interface{},
	opts ...options.Lister[options.UpdateManyOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateMany(context.Background(), filter, bson.M{
		"$set": data,
	}, opts...)
}
