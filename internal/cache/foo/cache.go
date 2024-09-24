package foo

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	redis *redis.Client
	ttl   time.Duration
}

func NewCache(redis *redis.Client, ttl time.Duration) *Cache {
	return &Cache{
		redis: redis,
		ttl:   ttl,
	}
}

func (c *Cache) generateKey(id string) string {
	return fmt.Sprintf("foo:%s", id)
}
