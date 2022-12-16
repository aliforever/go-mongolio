package mongolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) Distinct(filter bson.M, fieldName string, options ...*options.DistinctOptions) (result []interface{}, err error) {
	result, err = c.collection.Distinct(context.Background(), fieldName, filter, options...)
	if err != nil {
		return
	}

	return
}
