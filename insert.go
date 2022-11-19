package mongolio

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) InsertCustom(document any, options ...*options.InsertOneOptions) (result *mongo.InsertOneResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).InsertOne(context.Background(), document, options...)
	return
}

func (c *C[T]) Insert(document *T, options ...*options.InsertOneOptions) (result *mongo.InsertOneResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).InsertOne(context.Background(), document, options...)
	return
}

func (c *C[T]) InsertMany(documents []T, options ...*options.InsertManyOptions) (result *mongo.InsertManyResult, err error) {
	var interfaces []interface{}
	for _, document := range documents {
		interfaces = append(interfaces, document)
	}
	var i T
	result, err = c.db.Collection(i.CollectionName()).InsertMany(context.Background(), interfaces, options...)
	return
}
