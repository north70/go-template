package external_service

import (
	"context"
	"fmt"

	"github.com/north70/go-template/internal/domain"
	api "github.com/north70/go-template/internal/pb/external-service"
)

// GetFoo retrieves a Foo entity by its ID
func (c *Client) GetFoo(ctx context.Context, id string) (*domain.Foo, error) {
	req := mapGetFooRequest(id)
	resp, err := c.client.GetFoo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get foo: %w", err)
	}

	return mapGetFooResponse(resp), nil
}

func mapGetFooRequest(id string) *api.GetFooRequest {
	req := &api.GetFooRequest{Id: id}

	return req
}

func mapGetFooResponse(resp *api.Foo) *domain.Foo {
	return &domain.Foo{
		ID:   resp.GetId(),
		Name: resp.GetName(),
	}
}
