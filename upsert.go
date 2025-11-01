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
