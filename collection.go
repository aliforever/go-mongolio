package mongorm

import "go.mongodb.org/mongo-driver/mongo"

type col interface {
	Name() string
}

type collection[T col] struct {
	db *mongo.Database
}

func Collection[T col](db *mongo.Database) *collection[T] {
	return &collection[T]{
		db: db,
	}
}
