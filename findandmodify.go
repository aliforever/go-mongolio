package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// FindOneAndUpdate finds a document and updates it atomically, returns the document
func (c *C[T]) FindOneAndUpdate(
	filter bson.M,
	update *T,
	opts ...options.Lister[options.FindOneAndUpdateOptions],
) (*T, error) {
	var result T
	err := c.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		bson.M{"$set": update},
		opts...,
	).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindOneAndUpdateWithMap finds a document and updates specific fields atomically using a map
func (c *C[T]) FindOneAndUpdateWithMap(
	filter bson.M,
	update bson.M,
	opts ...options.Lister[options.FindOneAndUpdateOptions],
) (*T, error) {
	var result T
	err := c.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		bson.M{"$set": update},
		opts...,
	).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindOneAndUpdateOrdered finds a document and updates specific fields atomically using ordered bson.D
func (c *C[T]) FindOneAndUpdateOrdered(
	filter bson.M,
	update bson.D,
	opts ...options.Lister[options.FindOneAndUpdateOptions],
) (*T, error) {
	var result T
	err := c.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		bson.D{{"$set", update}},
		opts...,
	).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindOneAndUpdateCustom finds a document and updates it with custom operators
func (c *C[T]) FindOneAndUpdateCustom(
	filter bson.M,
	update *T,
	opts ...options.Lister[options.FindOneAndUpdateOptions],
) (*T, error) {
	var result T
	err := c.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		opts...,
	).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindOneAndUpdateCustomWithMap finds a document and updates it with custom operators using a map (no automatic $set wrapping)
func (c *C[T]) FindOneAndUpdateCustomWithMap(
	filter bson.M,
	update bson.M,
	opts ...options.Lister[options.FindOneAndUpdateOptions],
) (*T, error) {
	var result T
	err := c.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		opts...,
	).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindOneAndUpdateCustomOrdered finds a document and updates it with custom operators using ordered bson.D (no automatic $set wrapping)
func (c *C[T]) FindOneAndUpdateCustomOrdered(
	filter bson.M,
	update bson.D,
	opts ...options.Lister[options.FindOneAndUpdateOptions],
) (*T, error) {
	var result T
	err := c.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		opts...,
	).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindOneAndUpdateByID finds a document by ID and updates it atomically
func (c *C[T]) FindOneAndUpdateByID(
	id any,
	update *T,
	opts ...options.Lister[options.FindOneAndUpdateOptions],
) (*T, error) {
	return c.FindOneAndUpdate(bson.M{"_id": id}, update, opts...)
}

// FindOneAndReplace finds a document and replaces it atomically
func (c *C[T]) FindOneAndReplace(
	filter bson.M,
	replacement *T,
	opts ...options.Lister[options.FindOneAndReplaceOptions],
) (*T, error) {
	var result T
	err := c.collection.FindOneAndReplace(
		context.Background(),
		filter,
		replacement,
		opts...,
	).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindOneAndReplaceByID finds a document by ID and replaces it atomically
func (c *C[T]) FindOneAndReplaceByID(
	id any,
	replacement *T,
	opts ...options.Lister[options.FindOneAndReplaceOptions],
) (*T, error) {
	return c.FindOneAndReplace(bson.M{"_id": id}, replacement, opts...)
}

// FindOneAndDelete finds a document and deletes it atomically, returns the deleted document
func (c *C[T]) FindOneAndDelete(
	filter bson.M,
	opts ...options.Lister[options.FindOneAndDeleteOptions],
) (*T, error) {
	var result T
	err := c.collection.FindOneAndDelete(
		context.Background(),
		filter,
		opts...,
	).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindOneAndDeleteByID finds a document by ID and deletes it atomically
func (c *C[T]) FindOneAndDeleteByID(
	id any,
	opts ...options.Lister[options.FindOneAndDeleteOptions],
) (*T, error) {
	return c.FindOneAndDelete(bson.M{"_id": id}, opts...)
}
