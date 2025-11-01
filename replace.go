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
	replacement T,
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
	replacement T,
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
