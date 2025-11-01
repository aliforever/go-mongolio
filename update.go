package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) UpdateByID(
	id any,
	data *T,
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateByID(context.Background(), id, bson.M{"$set": data}, opts...)
}

// UpdateByIDWithMap updates specific fields of a document by ID using a map
func (c *C[T]) UpdateByIDWithMap(
	id any,
	data bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateByID(context.Background(), id, bson.M{"$set": data}, opts...)
}

// UpdateByIDOrdered updates specific fields of a document by ID using ordered bson.D
func (c *C[T]) UpdateByIDOrdered(
	id any,
	data bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateByID(context.Background(), id, bson.D{{"$set", data}}, opts...)
}

func (c *C[T]) UpdateCustom(
	filter bson.M,
	data *T,
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.M{
		"$set": data,
	}, opts...)
}

// UpdateCustomWithMap updates a document with custom operators using a map (no automatic $set wrapping)
func (c *C[T]) UpdateCustomWithMap(
	filter bson.M,
	update bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateOne(context.Background(), filter, update, opts...)
}

// UpdateCustomOrdered updates a document with custom operators using ordered bson.D (no automatic $set wrapping)
func (c *C[T]) UpdateCustomOrdered(
	filter bson.M,
	update bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateOne(context.Background(), filter, update, opts...)
}

// UpdateWithMap updates specific fields of a document using a map
func (c *C[T]) UpdateWithMap(
	filter bson.M,
	data bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.M{
		"$set": data,
	}, opts...)
}

// UpdateOrdered updates specific fields of a document using ordered bson.D
func (c *C[T]) UpdateOrdered(
	filter bson.M,
	data bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateOne(context.Background(), filter, bson.D{
		{"$set", data},
	}, opts...)
}

func (c *C[T]) UpdateCustomMany(
	filter bson.M,
	data *T,
	opts ...options.Lister[options.UpdateManyOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateMany(context.Background(), filter, bson.M{
		"$set": data,
	}, opts...)
}

// UpdateCustomManyWithMap updates multiple documents with custom operators using a map (no automatic $set wrapping)
func (c *C[T]) UpdateCustomManyWithMap(
	filter bson.M,
	update bson.M,
	opts ...options.Lister[options.UpdateManyOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateMany(context.Background(), filter, update, opts...)
}

// UpdateCustomManyOrdered updates multiple documents with custom operators using ordered bson.D (no automatic $set wrapping)
func (c *C[T]) UpdateCustomManyOrdered(
	filter bson.M,
	update bson.D,
	opts ...options.Lister[options.UpdateManyOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateMany(context.Background(), filter, update, opts...)
}

// UpdateManyWithMap updates specific fields of multiple documents using a map
func (c *C[T]) UpdateManyWithMap(
	filter bson.M,
	data bson.M,
	opts ...options.Lister[options.UpdateManyOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateMany(context.Background(), filter, bson.M{
		"$set": data,
	}, opts...)
}

// UpdateManyOrdered updates specific fields of multiple documents using ordered bson.D
func (c *C[T]) UpdateManyOrdered(
	filter bson.M,
	data bson.D,
	opts ...options.Lister[options.UpdateManyOptions],
) (result *mongo.UpdateResult, err error) {
	return c.collection.UpdateMany(context.Background(), filter, bson.D{
		{"$set", data},
	}, opts...)
}
