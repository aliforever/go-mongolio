package mongorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) Count(m bson.M, options ...*options.CountOptions) (result int64, err error) {
	var (
		i T
	)
	result, err = c.db.Collection(i.CollectionName()).CountDocuments(context.Background(), m, options...)
	return
}

func (c *C[T]) EstimatedCount(options ...*options.EstimatedDocumentCountOptions) (result int64, err error) {
	var (
		i T
	)
	result, err = c.db.Collection(i.CollectionName()).EstimatedDocumentCount(context.Background(), options...)
	return
}
