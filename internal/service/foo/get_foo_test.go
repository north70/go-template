package foo

import (
	"context"
	"testing"

	"github.com/north70/go-template/internal/cache"
	cacheMock "github.com/north70/go-template/internal/cache/mock"
	"github.com/north70/go-template/internal/domain"
	"github.com/north70/go-template/internal/repository"
	repoMock "github.com/north70/go-template/internal/repository/mock"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestService_GetFoo(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type fields struct {
		fooRepo  repository.FooRepository
		fooCache cache.FooCache
	}
	tests := []struct {
		name    string
		fields  fields
		id      string
		want    *domain.Foo
		wantErr error
	}{
		{
			name: "success, without cache",
			id:   "1",
			fields: fields{
				fooRepo: func() repository.FooRepository {
					fooRepoMock := repoMock.NewFooRepositoryMock(t)
					fooRepoMock.EXPECT().GetFoo(ctx, "1").
						Return(&domain.Foo{
							ID:   "1",
							Name: "Foo name",
						}, nil)

					return fooRepoMock
				}(),
				fooCache: func() cache.FooCache {
					fooCacheMock := cacheMock.NewFooCacheMock(t)
					fooCacheMock.EXPECT().Get(ctx, "1").
						Return(nil, redis.Nil)

					fooCacheMock.EXPECT().Set(ctx, &domain.Foo{
						ID:   "1",
						Name: "Foo name",
					}).Return(nil)

					return fooCacheMock
				}(),
			},
			want: &domain.Foo{
				ID:   "1",
				Name: "Foo name",
			},
			wantErr: nil,
		},
		{
			name: "success, with cache",
			id:   "1",
			fields: fields{
				fooCache: func() cache.FooCache {
					fooCacheMock := cacheMock.NewFooCacheMock(t)
					fooCacheMock.EXPECT().Get(ctx, "1").
						Return(&domain.Foo{
							ID:   "1",
							Name: "Foo name",
						}, nil)

					return fooCacheMock
				}(),
			},
			want: &domain.Foo{
				ID:   "1",
				Name: "Foo name",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Service{
				fooRepo:  tt.fields.fooRepo,
				fooCache: tt.fields.fooCache,
			}
			got, err := s.GetFoo(ctx, tt.id)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
