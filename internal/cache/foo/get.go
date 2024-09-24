package foo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/north70/go-template/internal/domain"

	"github.com/redis/go-redis/v9"
)

func (c *Cache) Get(ctx context.Context, id string) (*domain.Foo, error) {
	key := c.generateKey(id)

	value, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get foo from redis: %w", err)
	}

	foo := &domain.Foo{}
	if err := json.Unmarshal([]byte(value), foo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal foo: %w", err)
	}

	return foo, nil
}
