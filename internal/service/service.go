package service

import (
	"context"

	"github.com/north70/go-template/internal/domain"
)

//go:generate mockery --name FooService --structname FooServiceMock --outpkg mock --output ./mock --filename foo_service_mock.go --with-expecter
type FooService interface {
	GetFoo(ctx context.Context, id string) (*domain.Foo, error)
}
