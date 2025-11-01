package mongolio

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// PaginationResult contains paginated results and metadata
type PaginationResult[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	Page       int64 `json:"page"`
	PageSize   int64 `json:"page_size"`
	TotalPages int64 `json:"total_pages"`
}

// FindPaginated returns paginated results with metadata
func (c *C[T]) FindPaginated(
	filter bson.M,
	page int64,
	pageSize int64,
	opts ...options.Lister[options.FindOptions],
) (*PaginationResult[T], error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// Get total count
	total, err := c.Count(filter)
	if err != nil {
		return nil, err
	}

	// Calculate skip
	skip := (page - 1) * pageSize

	// Add pagination options
	findOpts := options.Find().SetSkip(skip).SetLimit(pageSize)
	allOpts := append([]options.Lister[options.FindOptions]{findOpts}, opts...)

	// Get paginated results
	results, err := c.Find(filter, allOpts...)
	if err != nil {
		return nil, err
	}

	// Calculate total pages
	totalPages := total / pageSize
	if total%pageSize != 0 {
		totalPages++
	}

	return &PaginationResult[T]{
		Data:       results,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// FindPaginatedOrdered returns paginated results with metadata using ordered bson.D filter
func (c *C[T]) FindPaginatedOrdered(
	filter bson.D,
	page int64,
	pageSize int64,
	opts ...options.Lister[options.FindOptions],
) (*PaginationResult[T], error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// Get total count
	total, err := c.CountOrdered(filter)
	if err != nil {
		return nil, err
	}

	// Calculate skip
	skip := (page - 1) * pageSize

	// Add pagination options
	findOpts := options.Find().SetSkip(skip).SetLimit(pageSize)
	allOpts := append([]options.Lister[options.FindOptions]{findOpts}, opts...)

	// Get paginated results
	results, err := c.FindOrdered(filter, allOpts...)
	if err != nil {
		return nil, err
	}

	// Calculate total pages
	totalPages := total / pageSize
	if total%pageSize != 0 {
		totalPages++
	}

	return &PaginationResult[T]{
		Data:       results,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// FindWithSkipLimit is a simple helper for skip/limit pagination
func (c *C[T]) FindWithSkipLimit(
	filter bson.M,
	skip int64,
	limit int64,
	opts ...options.Lister[options.FindOptions],
) ([]T, error) {
	findOpts := options.Find().SetSkip(skip).SetLimit(limit)
	allOpts := append([]options.Lister[options.FindOptions]{findOpts}, opts...)
	return c.Find(filter, allOpts...)
}

// FindWithSkipLimitOrdered is a simple helper for skip/limit pagination using ordered bson.D filter
func (c *C[T]) FindWithSkipLimitOrdered(
	filter bson.D,
	skip int64,
	limit int64,
	opts ...options.Lister[options.FindOptions],
) ([]T, error) {
	findOpts := options.Find().SetSkip(skip).SetLimit(limit)
	allOpts := append([]options.Lister[options.FindOptions]{findOpts}, opts...)
	return c.FindOrdered(filter, allOpts...)
}

// CountAndFind returns both count and results in a single call
func (c *C[T]) CountAndFind(
	filter bson.M,
	findOpts ...options.Lister[options.FindOptions],
) (int64, []T, error) {
	count, err := c.Count(filter)
	if err != nil {
		return 0, nil, err
	}

	results, err := c.Find(filter, findOpts...)
	if err != nil {
		return count, nil, err
	}

	return count, results, nil
}

// CountAndFindOrdered returns both count and results in a single call using ordered bson.D filter
func (c *C[T]) CountAndFindOrdered(
	filter bson.D,
	findOpts ...options.Lister[options.FindOptions],
) (int64, []T, error) {
	count, err := c.CountOrdered(filter)
	if err != nil {
		return 0, nil, err
	}

	results, err := c.FindOrdered(filter, findOpts...)
	if err != nil {
		return count, nil, err
	}

	return count, results, nil
}
