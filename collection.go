package mongorm

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type col interface {
	_CollectionName() string
}

type C[T col] struct {
	db *mongo.Database
}

func Collection[T col](db *mongo.Database) *C[T] {
	return &C[T]{
		db: db,
	}
}
