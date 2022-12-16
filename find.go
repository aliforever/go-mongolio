package mongolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) FindByID(id any, options ...*options.FindOneOptions) (result *T, err error) {
	err = c.collection.FindOne(context.Background(), bson.M{"_id": id}, options...).Decode(&result)
	return
}

func (c *C[T]) FindOne(m bson.M, options ...*options.FindOneOptions) (result *T, err error) {
	err = c.collection.FindOne(context.Background(), m, options...).Decode(&result)
	return
}

func (c *C[T]) Find(m bson.M, options ...*options.FindOptions) (result []T, err error) {
	var (
		cursor *mongo.Cursor
	)
	cursor, err = c.collection.Find(context.Background(), m, options...)
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
