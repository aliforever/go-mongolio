package mongorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) IncrementFieldByID(id any, fieldName string, amount int, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateByID(context.Background(), id, bson.M{"$inc": bson.M{fieldName: amount}}, options...)
	return
}

func (c *C[T]) IncrementFieldByFilter(filter bson.M, fieldName string, amount int, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateOne(context.Background(), filter, bson.M{"$inc": bson.M{fieldName: amount}}, options...)
	return
}

func (c *C[T]) IncrementFieldsByID(id any, fieldNameAmount bson.M, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateByID(context.Background(), id, bson.M{"$inc": fieldNameAmount}, options...)
	return
}

func (c *C[T]) IncrementFieldsByFilter(filter, fieldNameAmount bson.M, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateOne(context.Background(), filter, bson.M{"$inc": fieldNameAmount}, options...)
	return
}
