package cache

import (
	"context"

	"github.com/north70/go-template/internal/domain"
)

//go:generate mockery --name FooCache --structname FooCacheMock --outpkg mock --output ./mock --filename foo_cache_mock.go --with-expecter
type FooCache interface {
	Get(ctx context.Context, id string) (*domain.Foo, error)
	Set(ctx context.Context, foo *domain.Foo) error
}
