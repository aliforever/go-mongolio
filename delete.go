package mongolio

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) DeleteByID(id any, options ...*options.DeleteOptions) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteOne(context.Background(), bson.M{"_id": id}, options...)
	return
}

func (c *C[T]) DeleteOne(filter bson.M, options ...*options.DeleteOptions) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteOne(context.Background(), filter, options...)
	return
}

func (c *C[T]) DeleteMany(filter bson.M, options ...*options.DeleteOptions) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteMany(context.Background(), filter, options...)
	return
}
