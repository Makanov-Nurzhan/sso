package app

import (
	grpcapp "grpc-service/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: инициализировать хранилище (storage)
	// TODO: init auth service (auth)
	grpcServer := grpcapp.New(log, grpcPort)
	return &App{
		GRPCServer: grpcServer,
	}
}
