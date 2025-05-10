package grpc

import (
	"context"

	"github.com/lubie-koty/rpc-compute-service-combined/internal/core/types"
	pb "github.com/lubie-koty/rpc-compute-service-combined/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCService struct {
	pb.UnimplementedComplexComputeServer
	service types.MathService
}

func NewGRPCService(service types.MathService) *GRPCService {
	return &GRPCService{
		service: service,
	}
}

func (s *GRPCService) Sqrt(ctx context.Context, req *pb.UnaryRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}

func (s *GRPCService) Abs(ctx context.Context, req *pb.UnaryRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Abs not implemented")
}

func (s *GRPCService) Power(ctx context.Context, req *pb.BinaryRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Power not implemented")
}

func (s *GRPCService) Log(ctx context.Context, req *pb.BinaryRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}

func (s *GRPCService) Round(ctx context.Context, req *pb.BinaryRequest) (*pb.OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Round not implemented")
}
