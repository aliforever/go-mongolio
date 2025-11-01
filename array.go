package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// PushToArrayByID adds an item to an array field by document ID
func (c *C[T]) PushToArrayByID(
	id any,
	fieldName string,
	value any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$push": bson.M{fieldName: value}},
		opts...,
	)
}

// PushToArray adds an item to an array field
func (c *C[T]) PushToArray(
	filter bson.M,
	fieldName string,
	value any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$push": bson.M{fieldName: value}},
		opts...,
	)
}

// PushMultipleToArrayByID adds multiple items to an array field by document ID
func (c *C[T]) PushMultipleToArrayByID(
	id any,
	fieldName string,
	values []any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$push": bson.M{fieldName: bson.M{"$each": values}}},
		opts...,
	)
}

// PushMultipleToArray adds multiple items to an array field
func (c *C[T]) PushMultipleToArray(
	filter bson.M,
	fieldName string,
	values []any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$push": bson.M{fieldName: bson.M{"$each": values}}},
		opts...,
	)
}

// PullFromArrayByID removes items from an array field by document ID
func (c *C[T]) PullFromArrayByID(
	id any,
	fieldName string,
	value any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$pull": bson.M{fieldName: value}},
		opts...,
	)
}

// PullFromArray removes items from an array field
func (c *C[T]) PullFromArray(
	filter bson.M,
	fieldName string,
	value any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$pull": bson.M{fieldName: value}},
		opts...,
	)
}

// AddToSetByID adds unique items to an array field by document ID
func (c *C[T]) AddToSetByID(
	id any,
	fieldName string,
	value any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$addToSet": bson.M{fieldName: value}},
		opts...,
	)
}

// AddToSet adds unique items to an array field
func (c *C[T]) AddToSet(
	filter bson.M,
	fieldName string,
	value any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$addToSet": bson.M{fieldName: value}},
		opts...,
	)
}

// AddMultipleToSetByID adds multiple unique items to an array field by document ID
func (c *C[T]) AddMultipleToSetByID(
	id any,
	fieldName string,
	values []any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$addToSet": bson.M{fieldName: bson.M{"$each": values}}},
		opts...,
	)
}

// AddMultipleToSet adds multiple unique items to an array field
func (c *C[T]) AddMultipleToSet(
	filter bson.M,
	fieldName string,
	values []any,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$addToSet": bson.M{fieldName: bson.M{"$each": values}}},
		opts...,
	)
}

// PopFromArrayByID removes the first or last element from an array
// position: -1 for first element, 1 for last element
func (c *C[T]) PopFromArrayByID(
	id any,
	fieldName string,
	position int,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$pop": bson.M{fieldName: position}},
		opts...,
	)
}

// PopFromArray removes the first or last element from an array
// position: -1 for first element, 1 for last element
func (c *C[T]) PopFromArray(
	filter bson.M,
	fieldName string,
	position int,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$pop": bson.M{fieldName: position}},
		opts...,
	)
}
