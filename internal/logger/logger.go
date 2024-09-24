package logger

import (
	"context"
	"sync"
	"time"

	"github.com/north70/go-template/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
	once   sync.Once
)

// InitLogger initializes the global logger with custom configuration
func InitLogger(cfg config.Config) {
	once.Do(func() {
		level, err := zapcore.ParseLevel(cfg.App.LogLevel)
		if err != nil {
			level = zapcore.InfoLevel
		}

		config := zap.Config{
			Encoding:    "json",
			Level:       zap.NewAtomicLevelAt(level),
			OutputPaths: []string{"stdout"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:  "message",
				LevelKey:    "level",
				TimeKey:     "time",
				NameKey:     "logger",
				CallerKey:   "caller",
				LineEnding:  zapcore.DefaultLineEnding,
				EncodeLevel: zapcore.CapitalLevelEncoder,
				EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
					enc.AppendString(t.Format(time.RFC3339))
				},
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}

		zapLogger, _ := config.Build(zap.AddCallerSkip(1))
		logger = zapLogger.Sugar().With("service", cfg.App.ServiceName)
	})
}

// GetLogger returns the global SugaredLogger instance
func GetLogger() *zap.SugaredLogger {
	return logger
}

// SetLogLevel changes the logging level dynamically
func SetLogLevel(level string) {
	if l, err := zapcore.ParseLevel(level); err == nil {
		logger.Desugar().Core().Enabled(l)
	}
}

// Info logs a message at InfoLevel
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Error logs a message at ErrorLevel
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Debug logs a message at DebugLevel
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Warn logs a message at WarnLevel
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Fatal logs a message at FatalLevel
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// Infof logs a message at InfoLevel with format
func Infof(ctx context.Context, format string, args ...interface{}) {
	GetLogger().With("trace_id", ctx.Value("trace_id")).Infof(format, args...)
}

// Errorf logs a message at ErrorLevel with format
func Errorf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().With("trace_id", ctx.Value("trace_id")).Errorf(format, args...)
}

// Debugf logs a message at DebugLevel with format
func Debugf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().With("trace_id", ctx.Value("trace_id")).Debugf(format, args...)
}

// Warnf logs a message at WarnLevel with format
func Warnf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().With("trace_id", ctx.Value("trace_id")).Warnf(format, args...)
}

// Fatalf logs a message at FatalLevel with format
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	GetLogger().With("trace_id", ctx.Value("trace_id")).Fatalf(format, args...)
}
