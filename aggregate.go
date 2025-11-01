package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Aggregate runs an aggregation pipeline and returns typed results
func (c *C[T]) Aggregate(
	pipeline []bson.M,
	opts ...options.Lister[options.AggregateOptions],
) ([]T, error) {
	cursor, err := c.collection.Aggregate(context.Background(), pipeline, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []T
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// AggregateOne runs an aggregation pipeline and returns a single typed result
func (c *C[T]) AggregateOne(
	pipeline []bson.M,
	opts ...options.Lister[options.AggregateOptions],
) (*T, error) {
	cursor, err := c.collection.Aggregate(context.Background(), pipeline, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var result T
	if cursor.Next(context.Background()) {
		if err = cursor.Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return nil, mongo.ErrNoDocuments
}

// AggregateRaw runs an aggregation pipeline and returns raw bson documents
func (c *C[T]) AggregateRaw(
	pipeline []bson.M,
	opts ...options.Lister[options.AggregateOptions],
) ([]bson.M, error) {
	cursor, err := c.collection.Aggregate(context.Background(), pipeline, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
