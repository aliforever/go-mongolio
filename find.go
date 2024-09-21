package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) FindByID(id any, opts ...options.Lister[options.FindOneOptions]) (result *T, err error) {
	err = c.collection.FindOne(context.Background(), bson.M{"_id": id}, opts...).Decode(&result)
	return
}

func (c *C[T]) FindOne(m bson.M, opts ...options.Lister[options.FindOneOptions]) (result *T, err error) {
	err = c.collection.FindOne(context.Background(), m, opts...).Decode(&result)
	return
}

func (c *C[T]) Find(m bson.M, opts ...options.Lister[options.FindOptions]) (result []T, err error) {
	var (
		cursor *mongo.Cursor
	)
	cursor, err = c.collection.Find(context.Background(), m, opts...)
	if err != nil {
		return
	}

	err = cursor.All(context.Background(), &result)
	if err != nil {
		return
	}

	err = cursor.Err()
	return
}
