package mongorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *C[T]) Update(id primitive.ObjectID, document *T, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i._CollectionName()).UpdateByID(context.Background(), id, document, options...)
	return
}

func (c *C[T]) UpdateCustom(filter, data bson.M, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i._CollectionName()).UpdateOne(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
	return
}

func (c *C[T]) UpdateCustomMany(filter, data bson.M, options ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	var i T
	result, err = c.db.Collection(i._CollectionName()).UpdateMany(context.Background(), filter, bson.M{
		"$set": data,
	}, options...)
	return
}
