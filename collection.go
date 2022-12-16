package mongolio

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type C[T any] struct {
	collection *mongo.Collection
}

func Collection[T any](db *mongo.Database, collectionName string) *C[T] {
	return &C[T]{
		collection: db.Collection(collectionName),
	}
}
