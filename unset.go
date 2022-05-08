package mongorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *C[T]) UnsetFieldsByID(id any, fieldNames ...string) (result *mongo.UpdateResult, err error) {
	m := bson.M{}
	for _, name := range fieldNames {
		m[name] = ""
	}
	var i T
	result, err = c.db.Collection(i.CollectionName()).UpdateByID(context.Background(), id, bson.M{"$unset": m})
	return
}
