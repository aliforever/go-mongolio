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
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(context.Background(), id, bson.M{"$inc": bson.M{fieldName: amount}}, opts...)
}

func (c *C[T]) IncrementFieldByFilter(
	filter bson.M,
	fieldName string,
	amount int,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.M{"$inc": bson.M{fieldName: amount}}, opts...)
}

func (c *C[T]) IncrementFieldsByID(
	id any,
	fieldNameAmount bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(context.Background(), id, bson.M{"$inc": fieldNameAmount}, opts...)
}

// IncrementFieldsByIDOrdered increments multiple fields by ID using ordered bson.D
func (c *C[T]) IncrementFieldsByIDOrdered(
	id any,
	fieldNameAmount bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(context.Background(), id, bson.D{{"$inc", fieldNameAmount}}, opts...)
}

func (c *C[T]) IncrementFieldsByFilter(
	filter,
	fieldNameAmount bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.M{"$inc": fieldNameAmount}, opts...)
}

// IncrementFieldsByFilterOrdered increments multiple fields using ordered bson.D
func (c *C[T]) IncrementFieldsByFilterOrdered(
	filter bson.M,
	fieldNameAmount bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.D{{"$inc", fieldNameAmount}}, opts...)
}
