package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context, addr string, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return client, nil
}
