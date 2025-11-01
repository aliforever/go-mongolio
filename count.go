package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) Count(m bson.M, opts ...options.Lister[options.CountOptions]) (int64, error) {
	return c.collection.CountDocuments(context.Background(), m, opts...)
}

// CountOrdered counts documents using ordered bson.D filter
func (c *C[T]) CountOrdered(filter bson.D, opts ...options.Lister[options.CountOptions]) (int64, error) {
	return c.collection.CountDocuments(context.Background(), filter, opts...)
}

func (c *C[T]) EstimatedCount(opts ...options.Lister[options.EstimatedDocumentCountOptions]) (int64, error) {
	return c.collection.EstimatedDocumentCount(context.Background(), opts...)
}
