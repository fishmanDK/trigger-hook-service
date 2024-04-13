package app

import (
	grpcapp "github.com/fishmanDK/trigger_service/internal/app/grpc"
	"github.com/fishmanDK/trigger_service/internal/config"
	"github.com/fishmanDK/trigger_service/internal/service"
	"github.com/fishmanDK/trigger_service/internal/storage/postgres"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storageCFG config.PostgresConfig, tokenTTL time.Duration) *App {
	storage, err := postgres.NewPostgres(storageCFG)
	if err != nil {
		panic(err)
	}

	srvc := service.NewService(log, storage, tokenTTL)

	grpcApp := grpcapp.NewApp(log, srvc, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
