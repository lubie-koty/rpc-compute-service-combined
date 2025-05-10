package grpc

import (
	"context"
	"log/slog"
	"net"

	pb "github.com/lubie-koty/rpc-compute-service-combined/protos"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	Address           string
	Logger            *slog.Logger
	Service           *GRPCService
	Context           *context.Context
	ClientConnections []*grpc.ClientConn
}

func NewGRPCServer(ctx *context.Context, logger *slog.Logger, service *GRPCService, address string, connections []*grpc.ClientConn) *GRPCServer {
	return &GRPCServer{
		Address:           address,
		Logger:            logger,
		Service:           service,
		Context:           ctx,
		ClientConnections: connections,
	}
}

func (s *GRPCServer) Serve() error {
	lis, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterCombinedComputeServer(server, s.Service)

	s.Logger.Info("gRPC server started", "address", s.Address)
	g, gCtx := errgroup.WithContext(*s.Context)
	g.Go(func() error {
		return server.Serve(lis)
	})
	g.Go(func() error {
		<-gCtx.Done()
		s.Logger.Info("gRPC server stopped")
		for _, conn := range s.ClientConnections {
			conn.Close()
		}
		server.GracefulStop()
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
