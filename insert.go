package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *C[T]) InsertCustom(document any, opts ...options.Lister[options.InsertOneOptions]) (result *mongo.InsertOneResult, err error) {
	result, err = c.collection.InsertOne(context.Background(), document, opts...)
	return
}

func (c *C[T]) Insert(document *T, opts ...options.Lister[options.InsertOneOptions]) (result *mongo.InsertOneResult, err error) {
	result, err = c.collection.InsertOne(context.Background(), document, opts...)
	return
}

func (c *C[T]) InsertMany(documents []T, opts ...options.Lister[options.InsertManyOptions]) (result *mongo.InsertManyResult, err error) {
	var interfaces []interface{}
	for _, document := range documents {
		interfaces = append(interfaces, document)
	}
	result, err = c.collection.InsertMany(context.Background(), interfaces, opts...)
	return
}
