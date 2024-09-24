package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// RequestsTotal гистограмма общего количества запросов
	RequestsTotal = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "requests_total",
			Help:      "The total number of gRPC requests",
			Buckets:   prometheus.ExponentialBuckets(1, 2, 10),
		},
		[]string{"method", "status"},
	)

	// RequestDuration гистограмма времени выполнения запросов
	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "request_duration_seconds",
			Help:      "The duration of gRPC requests in seconds",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "status"},
	)

	// ErrorsTotal счетчик общего количества ошибок
	ErrorsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "errors_total",
			Help:      "The total number of gRPC errors",
		},
		[]string{"method", "status"},
	)
)
