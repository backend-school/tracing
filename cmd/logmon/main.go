package main

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

const (
	service = "demo-service"
	env     = "dev"
)

var (
	once   sync.Once
	logger *otelzap.Logger
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	tracer := otel.Tracer(service)

	ctx, span := tracer.Start(ctx, "demo")
	defer span.End()

	// Use Ctx to propagate the active span.
	logger := Logger(ctx)

	logger.Info("test info", zap.String("env", env))
	logger.Warn("test warn", zap.String("env", env))

	logger.Error("test error",
		zap.Error(errors.New("hello world")),
		zap.String("env", env))

	logger.Debug("test debug", zap.String("env", env))

	otelplay.PrintTraceID(ctx)
}

// Logger ensures that the caller does not forget to pass the context.
func Logger(ctx context.Context) otelzap.LoggerWithCtx {
	once.Do(func() {
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		logger = otelzap.New(l)
	})
	return logger.Ctx(ctx)
}
