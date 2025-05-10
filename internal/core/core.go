package core

import (
	"context"
	"log/slog"

	g "google.golang.org/grpc"

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
	simpleServiceURL := config.AppConfig.SimpleServiceURL
	complexServiceURL := config.AppConfig.ComplexServiceURL
	switch appMode {
	case "grpc":
		simpleClient, simpleConn := grpc.NewGRPCSimpleServiceClient(*ctx, logger, simpleServiceURL)
		complexClient, complexConn := grpc.NewGRPCComplexServiceClient(*ctx, logger, complexServiceURL)
		baseService := base.NewCombinedMathService(simpleClient, complexClient)
		server = grpc.NewGRPCServer(ctx, logger, grpc.NewGRPCService(baseService, logger), address, []*g.ClientConn{simpleConn, complexConn})
	case "rest":
		baseService := base.NewCombinedMathService(
			http.NewHTTPSimpleServiceClient(ctx, logger, simpleServiceURL),
			http.NewHTTPComplexServiceClient(ctx, logger, complexServiceURL),
		)
		server = http.NewHTTPServer(ctx, logger, http.NewHTTPService(baseService, logger), address)
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
