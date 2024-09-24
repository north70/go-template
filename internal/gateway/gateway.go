package gateway

import (
	"context"

	"github.com/north70/go-template/internal/domain"
)

type ExternalServiceClient interface {
	GetFoo(ctx context.Context, id string) (*domain.Foo, error)
}
