package external_service

import (
	api "github.com/north70/go-template/internal/pb/external-service"

	"google.golang.org/grpc"
)

// Client implements the ExternalServiceClient interface
type Client struct {
	client api.FooServiceClient
}

// NewClient creates a new Client instance
func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: api.NewFooServiceClient(conn),
	}
}
