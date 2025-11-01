package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Exists checks if any document matching the filter exists
func (c *C[T]) Exists(filter bson.M) (bool, error) {
	count, err := c.collection.CountDocuments(
		context.Background(),
		filter,
		options.Count().SetLimit(1),
	)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistsByID checks if a document with the given ID exists
func (c *C[T]) ExistsByID(id any) (bool, error) {
	return c.Exists(bson.M{"_id": id})
}

// ExistsWithProjection checks if a document exists and returns specific fields
func (c *C[T]) ExistsWithProjection(filter bson.M, projection bson.M) (*T, bool, error) {
	var result T
	err := c.collection.FindOne(
		context.Background(),
		filter,
		options.FindOne().SetProjection(projection),
	).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &result, true, nil
}

// CountGreaterThan checks if the count of documents matching filter is greater than n
func (c *C[T]) CountGreaterThan(filter bson.M, n int64) (bool, error) {
	count, err := c.collection.CountDocuments(
		context.Background(),
		filter,
		options.Count().SetLimit(n+1),
	)
	if err != nil {
		return false, err
	}
	return count > n, nil
}

// IsEmpty checks if the collection is empty
func (c *C[T]) IsEmpty() (bool, error) {
	count, err := c.collection.CountDocuments(
		context.Background(),
		bson.M{},
		options.Count().SetLimit(1),
	)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
