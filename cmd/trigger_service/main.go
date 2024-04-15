package main

import (
	"github.com/fishmanDK/trigger_service/internal/app"
	"github.com/fishmanDK/trigger_service/internal/config"
	"github.com/fishmanDK/trigger_service/internal/event_checker"
	"log/slog"
	"os"
	"time"
)

const (
	envLocal      = "local"
	envDev        = "dev"
	checkInterval = 3 * time.Minute
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(envLocal)

	log.Info("starting application", slog.Any("config", cfg))

	application := app.NewApp(log, cfg.GRPC.Port, cfg.Postgres, cfg.TokenTTL)

	log.Info("event_checker start")
	checker, err := event_checker.NewChecker(cfg.Postgres)
	if err != nil {
		log.Error("failed init checker", err)
		return
	}
	go func() {
		if err := checker.Run(1 * time.Minute); err != nil {
			log.Error("event_checker stopped", err)
		}
	}()

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
