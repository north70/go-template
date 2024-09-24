package repository

import (
	"context"
	"errors"

	"github.com/north70/go-template/internal/domain"
)

var (
	ErrNotFound = errors.New("not found")
)

//go:generate mockery --name FooRepository --structname FooRepositoryMock --outpkg mock --output ./mock --filename foo_repository_mock.go --with-expecter
type FooRepository interface {
	GetFoo(ctx context.Context, id string) (*domain.Foo, error)
}
