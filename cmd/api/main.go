package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/config"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/core"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/util"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	util.InitValidate()
	envVars, err := godotenv.Read()
	if err != nil {
		logger.Error("error loading .env file:", "error", err.Error())
		os.Exit(1)
	}
	err = config.InitConfig(envVars)
	if err != nil {
		logger.Error("error initializing config:", "error", err.Error())
		os.Exit(1)
	}

	err = runApplication(fmt.Sprintf("%s:%d", config.AppConfig.Host, config.AppConfig.Port), logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error("error running application", "error", err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func runApplication(address string, logger *slog.Logger) error {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		<-c
		cancel()
	}()

	app := core.NewApp(&ctx, logger, address)
	return app.Run()
}
