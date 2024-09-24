package api

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/north70/go-template/internal/domain"
	pb "github.com/north70/go-template/internal/pb/go-template"
	"github.com/north70/go-template/internal/repository"
	"github.com/north70/go-template/internal/service"
	"github.com/north70/go-template/internal/service/mock"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestApp_GetFoo(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name       string
		fooService service.FooService
		req        *pb.GetFooRequest
		want       *pb.Foo
		wantErr    error
	}{
		{
			name: "success",
			req: &pb.GetFooRequest{
				Id: "123",
			},
			fooService: func() service.FooService {
				fooServiceMock := mock.NewFooServiceMock(t)
				fooServiceMock.EXPECT().GetFoo(ctx, "123").
					Return(&domain.Foo{
						ID:   "123",
						Name: "test",
					}, nil)

				return fooServiceMock
			}(),
			want: &pb.Foo{
				Id:   "123",
				Name: "test",
			},
			wantErr: nil,
		},
		{
			name: "error: internal",
			req: &pb.GetFooRequest{
				Id: "123",
			},
			fooService: func() service.FooService {
				fooServiceMock := mock.NewFooServiceMock(t)
				fooServiceMock.EXPECT().GetFoo(ctx, "123").
					Return(nil, errors.New("internal"))

				return fooServiceMock
			}(),
			want:    nil,
			wantErr: status.Error(codes.Internal, "failed to get foo: internal"),
		},
		{
			name: "error: not found",
			req: &pb.GetFooRequest{
				Id: "123",
			},
			fooService: func() service.FooService {
				fooServiceMock := mock.NewFooServiceMock(t)
				fooServiceMock.EXPECT().GetFoo(ctx, "123").
					Return(nil, fmt.Errorf("not found: %w", repository.ErrNotFound))

				return fooServiceMock
			}(),
			want: nil,
			wantErr: status.Error(
				codes.NotFound,
				fmt.Sprintf("failed to get foo: not found: %v", repository.ErrNotFound)),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &App{fooService: tt.fooService}

			got, err := s.GetFoo(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
