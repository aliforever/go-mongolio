package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) BulkInsert(
	models []T,
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		wm = append(wm, mongo.NewInsertOneModel().SetDocument(model))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}
