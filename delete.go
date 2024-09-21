package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) DeleteByID(id any, opts ...options.Lister[options.DeleteOptions]) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteOne(context.Background(), bson.M{"_id": id}, opts...)
	return
}

func (c *C[T]) DeleteOne(filter bson.M, opts ...options.Lister[options.DeleteOptions]) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteOne(context.Background(), filter, opts...)
	return
}

func (c *C[T]) DeleteMany(filter bson.M, opts ...options.Lister[options.DeleteOptions]) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteMany(context.Background(), filter, opts...)
	return
}
