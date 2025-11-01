package mongolio

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// UpsertByID updates or inserts a document by ID
func (c *C[T]) UpsertByID(
	id any,
	data T,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	upsert := true
	updateOpts := options.UpdateOne().SetUpsert(upsert)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$set": data},
		allOpts...,
	)
}

// UpsertOne updates or inserts a document matching the filter
func (c *C[T]) UpsertOne(
	filter bson.M,
	data T,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	upsert := true
	updateOpts := options.UpdateOne().SetUpsert(upsert)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": data},
		allOpts...,
	)
}

// UpsertCustom updates or inserts with custom update operators
func (c *C[T]) UpsertCustom(
	filter bson.M,
	update T,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	upsert := true
	updateOpts := options.UpdateOne().SetUpsert(upsert)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		update,
		allOpts...,
	)
}

// BulkUpsert performs multiple upsert operations in a single bulk write
func (c *C[T]) BulkUpsert(
	models []UpdateModel[T],
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		wm = append(wm, mongo.NewUpdateOneModel().
			SetFilter(model.Filter).
			SetUpdate(bson.M{"$set": model.Model}).
			SetUpsert(true))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// BulkUpsertCustom performs multiple custom upsert operations in a single bulk write
func (c *C[T]) BulkUpsertCustom(
	models []UpdateModel[T],
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		wm = append(wm, mongo.NewUpdateOneModel().
			SetFilter(model.Filter).
			SetUpdate(model.Model).
			SetUpsert(true))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}
