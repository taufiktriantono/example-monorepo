package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals"
	"github.com/taufiktriantono/api-first-monorepo/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var fxLogger = fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log}
})

var zapField = fx.Provide(func(cfg *config.Config) []zap.Field {
	return []zap.Field{
		zap.String("service_name", cfg.AppName),
		zap.String("environment", cfg.AppEnv),
	}
})

var engine = fx.Provide(func() *gin.Engine {
	return gin.New()
})

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := fx.New(
		engine,
		zapField,
		approvals.Service,
		fxLogger,
	)

	if err := app.Start(ctx); err != nil {
		zap.L().Fatal("Failed to start", zap.Error(err))
	}

	// Wait for signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	zap.L().Info("Server is shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := app.Stop(shutdownCtx); err != nil {
		zap.L().Error("Failed to stop", zap.Error(err))
	}
}
