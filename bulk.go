package mongolio

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) BulkImport(models []mongo.WriteModel, options ...*options.BulkWriteOptions) (result *mongo.BulkWriteResult, err error) {
	var i T

	result, err = c.db.Collection(i.CollectionName()).BulkWrite(context.Background(), models, options...)
	return
}
