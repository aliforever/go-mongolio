package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// BulkDeleteByIDs deletes multiple documents by their IDs
func (c *C[T]) BulkDeleteByIDs(
	ids []any,
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, id := range ids {
		wm = append(wm, mongo.NewDeleteOneModel().
			SetFilter(bson.M{"_id": id}))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// BulkDelete performs multiple delete operations based on filters
func (c *C[T]) BulkDelete(
	filters []bson.M,
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, filter := range filters {
		wm = append(wm, mongo.NewDeleteOneModel().
			SetFilter(filter))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// BulkDeleteMany performs multiple deleteMany operations based on filters
func (c *C[T]) BulkDeleteMany(
	filters []bson.M,
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, filter := range filters {
		wm = append(wm, mongo.NewDeleteManyModel().
			SetFilter(filter))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}
