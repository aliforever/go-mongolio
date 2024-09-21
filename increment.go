package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) IncrementFieldByID(
	id any,
	fieldName string,
	amount int,
	opts ...options.Lister[options.UpdateOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(context.Background(), id, bson.M{"$inc": bson.M{fieldName: amount}}, opts...)
}

func (c *C[T]) IncrementFieldByFilter(
	filter bson.M,
	fieldName string,
	amount int,
	opts ...options.Lister[options.UpdateOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.M{"$inc": bson.M{fieldName: amount}}, opts...)
}

func (c *C[T]) IncrementFieldsByID(
	id any,
	fieldNameAmount bson.M,
	opts ...options.Lister[options.UpdateOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(context.Background(), id, bson.M{"$inc": fieldNameAmount}, opts...)
}

func (c *C[T]) IncrementFieldsByFilter(
	filter,
	fieldNameAmount bson.M,
	opts ...options.Lister[options.UpdateOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.M{"$inc": fieldNameAmount}, opts...)
}
