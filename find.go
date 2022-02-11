package mongorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *collection[T]) FindOne(m bson.M, options ...*options.FindOneOptions) (result *T, err error) {
	var i T
	err = c.db.Collection(i.Name()).FindOne(context.Background(), m, options...).Decode(&result)
	return
}
