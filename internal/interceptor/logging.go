package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// LoggingInterceptor ...
func LoggingInterceptor(log *zap.SugaredLogger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		duration := time.Since(start)

		md, _ := metadata.FromIncomingContext(ctx)
		traceID := getTraceID(md)
		s, _ := status.FromError(err)

		msg := fmt.Sprintf("'%s' responded: '%s' in %d ms", info.FullMethod, s.Code(), duration.Milliseconds())

		switch s.Code() {
		case codes.OK:
			if log.Level() == zap.DebugLevel {
				log.Debug(zap.String("message", msg),
					zap.String("method", info.FullMethod),
					zap.Any("request", req),
					zap.Any("response", resp),
					zap.String("trace_id", traceID),
				)
			}

			log.Info(
				zap.String("message", msg),
				zap.String("method", info.FullMethod),
				zap.Any("request", req),
			)
		default:
			log.Error(
				zap.String("message", msg),
				zap.String("method", info.FullMethod),
				zap.Any("request", resp),
				zap.Any("response", s.Message()),
				zap.String("trace_id", traceID),
			)
		}

		return resp, err
	}
}

func getTraceID(md metadata.MD) string {
	if len(md["x-trace-id"]) > 0 {
		return md["x-trace-id"][0]
	}

	return uuid.New().String()
}
