package foo

import (
	"context"
	"errors"
	"fmt"

	"github.com/north70/go-template/internal/domain"
	"github.com/north70/go-template/internal/logger"

	"github.com/redis/go-redis/v9"
)

func (s *Service) GetFoo(ctx context.Context, id string) (*domain.Foo, error) {
	foo, err := s.fooCache.Get(ctx, id)
	if err == nil && foo != nil {
		return foo, nil
	}

	if err != nil && !errors.Is(err, redis.Nil) {
		logger.Warnf(ctx, "failed to get foo from cache: %v", err)
	}

	foo, err = s.fooRepo.GetFoo(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get foo from repository: %w", err)
	}

	if err = s.fooCache.Set(ctx, foo); err != nil {
		logger.Warnf(ctx, "failed to set foo in cache: %v", err)
	}

	return foo, nil
}
