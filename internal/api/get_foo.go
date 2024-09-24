package api

import (
	"context"

	pb "github.com/north70/go-template/internal/pb/go-template"

	"google.golang.org/grpc/status"
)

func (s *App) GetFoo(ctx context.Context, req *pb.GetFooRequest) (*pb.Foo, error) {
	foo, err := s.fooService.GetFoo(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(parseErrorCodes(err), "failed to get foo: %v", err)
	}

	return &pb.Foo{
		Id:    foo.ID,
		Name:  foo.Name,
		Value: 1234,
	}, nil
}
