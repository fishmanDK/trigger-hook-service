package app

import (
	grpcapp "github.com/fishmanDK/trigger_service/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
