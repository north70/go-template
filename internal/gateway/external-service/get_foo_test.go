package external_service

import (
	"testing"

	"github.com/north70/go-template/internal/domain"
	api "github.com/north70/go-template/internal/pb/external-service"

	"github.com/stretchr/testify/assert"
)

func Test_mapGetFooRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		id   string
		want *api.GetFooRequest
	}{
		{
			name: "success",
			id:   "1",
			want: &api.GetFooRequest{
				Id: "1",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := mapGetFooRequest(tt.id)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_mapGetFooResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		resp *api.Foo
		want *domain.Foo
	}{
		{
			name: "success",
			resp: &api.Foo{
				Id:    "1",
				Name:  "test",
				Value: 123,
			},
			want: &domain.Foo{
				ID:   "1",
				Name: "test",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := mapGetFooResponse(tt.resp)
			assert.Equal(t, tt.want, got)
		})
	}
}
