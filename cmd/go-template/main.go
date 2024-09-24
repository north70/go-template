package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/north70/go-template/internal/api"
	fooCache "github.com/north70/go-template/internal/cache/foo"
	"github.com/north70/go-template/internal/config"
	external_service "github.com/north70/go-template/internal/gateway/external-service"
	"github.com/north70/go-template/internal/interceptor"
	"github.com/north70/go-template/internal/logger"
	pb "github.com/north70/go-template/internal/pb/go-template"
	fooRepository "github.com/north70/go-template/internal/repository/foo"
	fooService "github.com/north70/go-template/internal/service/foo"
	"github.com/north70/go-template/internal/storage"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		fmt.Println("Config file path is required")
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger.InitLogger(*cfg)
	log := logger.GetLogger()

	grpcServer := startGrpcServer(ctx, cfg, log)
	httpServer := startHttpServer(ctx, cfg, log)
	metricsServer := startMetricsServer(ctx, cfg, log)

	gracefulShutdown(ctx, log, grpcServer, httpServer, metricsServer)
}

func startHttpServer(ctx context.Context, cfg *config.Config, log *zap.SugaredLogger) *http.Server {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterFooServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", cfg.App.GRPCPort), opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway: %v", err)
	}

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.App.HTTPPort),
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Infof("Starting HTTP server on port %d", cfg.App.HTTPPort)
		if err = httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()
	return httpServer
}

func startMetricsServer(ctx context.Context, cfg *config.Config, log *zap.SugaredLogger) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	metricsServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.App.MetricsPort),
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Infof("Starting metrics server on port %d", cfg.App.MetricsPort)
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start metrics server: %v", err)
		}
	}()

	return metricsServer
}

func startGrpcServer(ctx context.Context, cfg *config.Config, log *zap.SugaredLogger) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.LoggingInterceptor(log),
			interceptor.MetricsInterceptor(),
		),
	)
	app, err := initApp(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	pb.RegisterFooServiceServer(grpcServer, app)

	reflection.Register(grpcServer)

	go func() {
		log.Infof("Starting gRPC server on port %d", cfg.App.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	return grpcServer
}

func gracefulShutdown(
	ctx context.Context,
	log *zap.SugaredLogger,
	grpcServer *grpc.Server,
	httpServer *http.Server,
	metricsServer *http.Server,
) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down servers...")

	grpcServer.GracefulStop()

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server forced to shutdown: %v", err)
	}

	if err := metricsServer.Shutdown(ctx); err != nil {
		log.Fatalf("Metrics server forced to shutdown: %v", err)
	}

	log.Info("Servers exiting")
}

func initApp(ctx context.Context, cfg *config.Config) (*api.App, error) {
	postgres, err := storage.NewPostgres(ctx, cfg.Database.PostgresDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres: %w", err)
	}

	redis, err := storage.NewRedis(ctx, cfg.Database.RedisAddr, cfg.Database.RedisPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to create redis: %w", err)
	}

	repoFoo := fooRepository.NewRepository(postgres)

	cacheFoo := fooCache.NewCache(redis, time.Minute)

	conn, err := grpc.NewClient(cfg.Client.ExternalServiceEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create external service client: %w", err)
	}
	externalServiceClient := external_service.NewClient(conn)

	serviceFoo := fooService.NewService(repoFoo, cacheFoo, externalServiceClient)

	return api.NewApp(serviceFoo), nil
}
