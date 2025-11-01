package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) Distinct(
	filter bson.M,
	fieldName string,
	opts ...options.Lister[options.DistinctOptions],
) ([]T, error) {
	var data []T

	if err := c.collection.Distinct(context.Background(), fieldName, filter, opts...).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DistinctOrdered gets distinct values using ordered bson.D filter
func (c *C[T]) DistinctOrdered(
	filter bson.D,
	fieldName string,
	opts ...options.Lister[options.DistinctOptions],
) ([]T, error) {
	var data []T

	if err := c.collection.Distinct(context.Background(), fieldName, filter, opts...).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
