package main

import (
	"github.com/fishmanDK/trigger_service/internal/app"
	"github.com/fishmanDK/trigger_service/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(envLocal)

	log.Info("starting application", slog.Any("config", cfg))

	application := app.NewApp(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.GRPCSrv.Run()
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		slogHandler := slog.NewTextHandler(os.Stdout, opts)
		logger = slog.New(slogHandler)
	case envDev:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		slogHandler := slog.NewJSONHandler(os.Stdout, opts)
		logger = slog.New(slogHandler)
	}

	return logger
}
