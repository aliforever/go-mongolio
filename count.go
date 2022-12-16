package mongolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) Count(m bson.M, options ...*options.CountOptions) (result int64, err error) {
	result, err = c.collection.CountDocuments(context.Background(), m, options...)
	return
}

func (c *C[T]) EstimatedCount(options ...*options.EstimatedDocumentCountOptions) (result int64, err error) {
	result, err = c.collection.EstimatedDocumentCount(context.Background(), options...)
	return
}
