package foo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/north70/go-template/internal/domain"
)

func (c *Cache) Set(ctx context.Context, foo *domain.Foo) error {
	key := c.generateKey(foo.ID)

	value, err := json.Marshal(foo)
	if err != nil {
		return fmt.Errorf("failed to marshal foo: %w", err)
	}

	return c.redis.Set(ctx, key, value, c.ttl).Err()
}
