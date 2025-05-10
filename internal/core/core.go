package core

import (
	"context"
	"log/slog"

	"github.com/lubie-koty/rpc-compute-service-combined/internal/base"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/config"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/core/types"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/grpc"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/http"
)

type App struct {
	Server types.Server
}

func NewApp(ctx *context.Context, logger *slog.Logger, address string) *App {
	var server types.Server
	appMode := config.AppConfig.AppMode
	baseService := base.NewCombinedMathService()
	switch appMode {
	case "grpc":
		server = grpc.NewGRPCServer(ctx, logger, grpc.NewGRPCService(baseService), address)
	case "rest":
		server = http.NewHTTPServer(ctx, logger, http.NewHTTPService(baseService), address)
	default:
		panic("unsupported app mode")
	}
	return &App{
		Server: server,
	}
}

func (a *App) Run() error {
	return a.Server.Serve()
}
