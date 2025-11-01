package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) DeleteByID(id any, opts ...options.Lister[options.DeleteOneOptions]) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteOne(context.Background(), bson.M{"_id": id}, opts...)
	return
}

func (c *C[T]) DeleteOne(filter bson.M, opts ...options.Lister[options.DeleteOneOptions]) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteOne(context.Background(), filter, opts...)
	return
}

// DeleteOneOrdered deletes a single document using ordered bson.D filter
func (c *C[T]) DeleteOneOrdered(filter bson.D, opts ...options.Lister[options.DeleteOneOptions]) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteOne(context.Background(), filter, opts...)
	return
}

func (c *C[T]) DeleteMany(filter bson.M, opts ...options.Lister[options.DeleteManyOptions]) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteMany(context.Background(), filter, opts...)
	return
}

// DeleteManyOrdered deletes multiple documents using ordered bson.D filter
func (c *C[T]) DeleteManyOrdered(filter bson.D, opts ...options.Lister[options.DeleteManyOptions]) (result *mongo.DeleteResult, err error) {
	result, err = c.collection.DeleteMany(context.Background(), filter, opts...)
	return
}
