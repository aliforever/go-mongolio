package mongolio

import (
	"context"
)

func (c *C[T]) Drop() (err error) {
	return c.collection.Drop(context.Background())
}
