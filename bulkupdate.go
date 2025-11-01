package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// UpdateModel represents a single update operation with generic update data type
type UpdateModel[T any] struct {
	Filter bson.M
	Model  T
}

// BulkUpdate performs multiple update operations in a single bulk write
func (c *C[T]) BulkUpdate(
	models []UpdateModel[T],
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		wm = append(wm, mongo.NewUpdateOneModel().
			SetFilter(model.Filter).
			SetUpdate(bson.M{"$set": model.Model}))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// BulkUpdateMany performs multiple updateMany operations in a single bulk write
func (c *C[T]) BulkUpdateMany(
	models []UpdateModel[T],
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		wm = append(wm, mongo.NewUpdateManyModel().
			SetFilter(model.Filter).
			SetUpdate(bson.M{"$set": model.Model}))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// BulkUpdateByID performs multiple update operations by ID in a single bulk write
func (c *C[T]) BulkUpdateByID(
	updates map[any]T,
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for id, data := range updates {
		wm = append(wm, mongo.NewUpdateOneModel().
			SetFilter(bson.M{"_id": id}).
			SetUpdate(bson.M{"$set": data}))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// BulkUpdateCustom performs multiple custom update operations (with custom update operators)
func (c *C[T]) BulkUpdateCustom(
	models []UpdateModel[T],
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		wm = append(wm, mongo.NewUpdateOneModel().
			SetFilter(model.Filter).
			SetUpdate(bson.M{"$set": model.Model}))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// BulkUpdateCustomMany performs multiple custom updateMany operations (with custom update operators)
func (c *C[T]) BulkUpdateCustomMany(
	models []UpdateModel[T],
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		wm = append(wm, mongo.NewUpdateManyModel().
			SetFilter(model.Filter).
			SetUpdate(bson.M{"$set": model.Model}))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}
