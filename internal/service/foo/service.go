package foo

import (
	"github.com/north70/go-template/internal/cache"
	"github.com/north70/go-template/internal/gateway"
	"github.com/north70/go-template/internal/repository"
)

type Service struct {
	fooRepo        repository.FooRepository
	fooCache       cache.FooCache
	externalClient gateway.ExternalServiceClient
}

func NewService(
	fooRepo repository.FooRepository,
	fooCache cache.FooCache,
	externalClient gateway.ExternalServiceClient,
) *Service {
	return &Service{
		fooRepo:        fooRepo,
		fooCache:       fooCache,
		externalClient: externalClient,
	}
}
