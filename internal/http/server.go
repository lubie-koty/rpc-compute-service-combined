package http

import (
	"context"
	"log/slog"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type HTTPServer struct {
	Address string
	Logger  *slog.Logger
	Context *context.Context
	*HTTPService
}

func NewHTTPServer(ctx *context.Context, logger *slog.Logger, httpService *HTTPService, address string) *HTTPServer {
	return &HTTPServer{
		Address:     address,
		Logger:      logger,
		Context:     ctx,
		HTTPService: httpService,
	}
}

func NewHTTPHandler(httpService *HTTPService) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /rms", httpService.RootMeanSquare)
	mux.HandleFunc("GET /gm", httpService.GeometricMean)
	mux.HandleFunc("GET /bmi", httpService.BodyMassIndex)
	mux.HandleFunc("POST /pld", httpService.PowerLevelDiff)
	mux.HandleFunc("POST /pvc", httpService.PercentageValueChange)
	return mux
}

func (s *HTTPServer) Serve() error {
	server := &http.Server{
		Addr:    s.Address,
		Handler: NewHTTPHandler(s.HTTPService),
	}

	s.Logger.Info("HTTP server started", "server address", server.Addr)
	g, gCtx := errgroup.WithContext(*s.Context)
	g.Go(func() error {
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		s.Logger.Info("HTTP server stopped")
		return server.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
