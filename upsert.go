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
	data *T,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$set": data},
		allOpts...,
	)
}

// UpsertByIDWithMap updates or inserts specific fields by ID using a map
func (c *C[T]) UpsertByIDWithMap(
	id any,
	data bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.M{"$set": data},
		allOpts...,
	)
}

// UpsertByIDOrdered updates or inserts specific fields by ID using ordered bson.D
func (c *C[T]) UpsertByIDOrdered(
	id any,
	data bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateByID(
		context.Background(),
		id,
		bson.D{{"$set", data}},
		allOpts...,
	)
}

// UpsertOne updates or inserts a document matching the filter
func (c *C[T]) UpsertOne(
	filter bson.M,
	data *T,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": data},
		allOpts...,
	)
}

// UpsertWithMap updates or inserts specific fields using a map
func (c *C[T]) UpsertWithMap(
	filter bson.M,
	data bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": data},
		allOpts...,
	)
}

// UpsertOrdered updates or inserts specific fields using ordered bson.D
func (c *C[T]) UpsertOrdered(
	filter bson.M,
	data bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.D{{"$set", data}},
		allOpts...,
	)
}

// UpsertCustom updates or inserts with custom update operators
func (c *C[T]) UpsertCustom(
	filter bson.M,
	update *T,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": update},
		allOpts...,
	)
}

// UpsertCustomWithMap updates or inserts with custom update operators using a map
func (c *C[T]) UpsertCustomWithMap(
	filter bson.M,
	update bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.M{"$set": update},
		allOpts...,
	)
}

// UpsertCustomOrdered updates or inserts with custom update operators using ordered bson.D
func (c *C[T]) UpsertCustomOrdered(
	filter bson.M,
	update bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		bson.D{{"$set", update}},
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
			SetUpdate(bson.M{"$set": model.Model}).
			SetUpsert(true))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// UpsertByIDWithMapAndSetOnInsert updates or inserts specific fields by ID with separate $set and $setOnInsert maps
func (c *C[T]) UpsertByIDWithMapAndSetOnInsert(
	id any,
	setData bson.M,
	setOnInsertData bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	update := bson.M{
		"$set":         setData,
		"$setOnInsert": setOnInsertData,
	}

	return c.collection.UpdateByID(
		context.Background(),
		id,
		update,
		allOpts...,
	)
}

// UpsertByIDOrderedWithSetOnInsert updates or inserts specific fields by ID with separate $set and $setOnInsert using ordered bson.D
func (c *C[T]) UpsertByIDOrderedWithSetOnInsert(
	id any,
	setData bson.D,
	setOnInsertData bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	update := bson.D{
		{"$set", setData},
		{"$setOnInsert", setOnInsertData},
	}

	return c.collection.UpdateByID(
		context.Background(),
		id,
		update,
		allOpts...,
	)
}

// UpsertWithMapAndSetOnInsert updates or inserts specific fields with separate $set and $setOnInsert maps
func (c *C[T]) UpsertWithMapAndSetOnInsert(
	filter bson.M,
	setData bson.M,
	setOnInsertData bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	update := bson.M{
		"$set":         setData,
		"$setOnInsert": setOnInsertData,
	}

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		update,
		allOpts...,
	)
}

// UpsertOrderedWithSetOnInsert updates or inserts specific fields with separate $set and $setOnInsert using ordered bson.D
func (c *C[T]) UpsertOrderedWithSetOnInsert(
	filter bson.M,
	setData bson.D,
	setOnInsertData bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	update := bson.D{
		{"$set", setData},
		{"$setOnInsert", setOnInsertData},
	}

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		update,
		allOpts...,
	)
}

// UpsertCustomWithMapAndSetOnInsert updates or inserts with custom update operators using maps for both $set and $setOnInsert
func (c *C[T]) UpsertCustomWithMapAndSetOnInsert(
	filter bson.M,
	setData bson.M,
	setOnInsertData bson.M,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	update := bson.M{
		"$set":         setData,
		"$setOnInsert": setOnInsertData,
	}

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		update,
		allOpts...,
	)
}

// UpsertCustomOrderedWithSetOnInsert updates or inserts with custom update operators using ordered bson.D for both $set and $setOnInsert
func (c *C[T]) UpsertCustomOrderedWithSetOnInsert(
	filter bson.M,
	setData bson.D,
	setOnInsertData bson.D,
	opts ...options.Lister[options.UpdateOneOptions],
) (*mongo.UpdateResult, error) {
	updateOpts := options.UpdateOne().SetUpsert(true)
	allOpts := append([]options.Lister[options.UpdateOneOptions]{updateOpts}, opts...)

	update := bson.D{
		{"$set", setData},
		{"$setOnInsert", setOnInsertData},
	}

	return c.collection.UpdateOne(
		context.Background(),
		filter,
		update,
		allOpts...,
	)
}

// UpdateModelWithSetOnInsert represents a single update operation with separate $set and $setOnInsert data using maps
type UpdateModelWithSetOnInsert struct {
	Filter          bson.M
	SetData         bson.M
	SetOnInsertData bson.M
}

// BulkUpsertWithSetOnInsert performs multiple upsert operations with $setOnInsert in a single bulk write
func (c *C[T]) BulkUpsertWithSetOnInsert(
	models []UpdateModelWithSetOnInsert,
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		update := bson.M{
			"$set":         model.SetData,
			"$setOnInsert": model.SetOnInsertData,
		}
		wm = append(wm, mongo.NewUpdateOneModel().
			SetFilter(model.Filter).
			SetUpdate(update).
			SetUpsert(true))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}

// BulkUpsertCustomWithSetOnInsert performs multiple custom upsert operations with $setOnInsert in a single bulk write
func (c *C[T]) BulkUpsertCustomWithSetOnInsert(
	models []UpdateModelWithSetOnInsert,
	opts ...options.Lister[options.BulkWriteOptions],
) (*mongo.BulkWriteResult, error) {
	var wm []mongo.WriteModel

	for _, model := range models {
		update := bson.M{
			"$set":         model.SetData,
			"$setOnInsert": model.SetOnInsertData,
		}
		wm = append(wm, mongo.NewUpdateOneModel().
			SetFilter(model.Filter).
			SetUpdate(update).
			SetUpsert(true))
	}

	return c.collection.BulkWrite(context.Background(), wm, opts...)
}
