package mongorm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type col interface {
	Name() string
	ID() primitive.ObjectID
}

type C[T col] struct {
	db *mongo.Database
}

func Collection[T col](db *mongo.Database) *C[T] {
	return &C[T]{
		db: db,
	}
}
