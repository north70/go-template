package api

import (
	"errors"

	pb "github.com/north70/go-template/internal/pb/go-template"
	"github.com/north70/go-template/internal/repository"
	"github.com/north70/go-template/internal/service"

	"google.golang.org/grpc/codes"
)

type App struct {
	pb.UnimplementedFooServiceServer
	fooService service.FooService
}

func NewApp(fooService service.FooService) *App {
	return &App{fooService: fooService}
}

func parseErrorCodes(err error) codes.Code {
	if errors.Is(err, repository.ErrNotFound) {
		return codes.NotFound
	}
	return codes.Internal
}
