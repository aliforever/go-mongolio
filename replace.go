package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ReplaceByID replaces an entire document by ID
func (c *C[T]) ReplaceByID(
	id any,
	replacement *T,
	opts ...options.Lister[options.ReplaceOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.ReplaceOne(
		context.Background(),
		bson.M{"_id": id},
		replacement,
		opts...,
	)
}

// ReplaceByIDWithMap replaces an entire document by ID using a map
func (c *C[T]) ReplaceByIDWithMap(
	id any,
	replacement bson.M,
	opts ...options.Lister[options.ReplaceOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.ReplaceOne(
		context.Background(),
		bson.M{"_id": id},
		replacement,
		opts...,
	)
}

// ReplaceByIDOrdered replaces an entire document by ID using ordered bson.D
func (c *C[T]) ReplaceByIDOrdered(
	id any,
	replacement bson.D,
	opts ...options.Lister[options.ReplaceOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.ReplaceOne(
		context.Background(),
		bson.M{"_id": id},
		replacement,
		opts...,
	)
}

// ReplaceOne replaces an entire document matching the filter
func (c *C[T]) ReplaceOne(
	filter bson.M,
	replacement *T,
	opts ...options.Lister[options.ReplaceOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.ReplaceOne(
		context.Background(),
		filter,
		replacement,
		opts...,
	)
}

// ReplaceOneWithMap replaces an entire document using a map
func (c *C[T]) ReplaceOneWithMap(
	filter bson.M,
	replacement bson.M,
	opts ...options.Lister[options.ReplaceOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.ReplaceOne(
		context.Background(),
		filter,
		replacement,
		opts...,
	)
}

// ReplaceOneOrdered replaces an entire document using ordered bson.D
func (c *C[T]) ReplaceOneOrdered(
	filter bson.M,
	replacement bson.D,
	opts ...options.Lister[options.ReplaceOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.ReplaceOne(
		context.Background(),
		filter,
		replacement,
		opts...,
	)
}

// BulkReplace performs multiple replace operations in a single bulk write
func (c *C[T]) BulkReplace(
	models []UpdateModel[T],
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		wm = append(wm, mongo.NewReplaceOneModel().
			SetFilter(model.Filter).
			SetReplacement(model.Model))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}
