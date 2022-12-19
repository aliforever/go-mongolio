package mongolio

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

	return c.collection.UpdateByID(context.Background(), id, bson.M{"$unset": m})
}
