package interceptor

import (
	"context"
	"time"

	"github.com/north70/go-template/internal/metrics"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// MetricsInterceptor создает новый gRPC сервер интерцептор для сбора метрик
func MetricsInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		resp, err := handler(ctx, req)

		statusCode := status.Code(err)

		duration := time.Since(startTime).Seconds()

		metrics.RequestsTotal.WithLabelValues(info.FullMethod, statusCode.String()).Observe(1)

		metrics.RequestDuration.WithLabelValues(info.FullMethod, statusCode.String()).Observe(duration)

		if err != nil {
			metrics.ErrorsTotal.WithLabelValues(info.FullMethod, statusCode.String()).Inc()
		}

		return resp, err
	}
}
