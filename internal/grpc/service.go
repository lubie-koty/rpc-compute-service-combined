package grpc

import (
	"context"

	"github.com/lubie-koty/rpc-compute-service-combined/internal/core/types"
	pb "github.com/lubie-koty/rpc-compute-service-combined/protos"
)

type GRPCService struct {
	pb.UnimplementedCombinedComputeServer
	service types.MathService
}

func NewGRPCService(service types.MathService) *GRPCService {
	return &GRPCService{
		service: service,
	}
}

func (s *GRPCService) RootMeanSquare(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := s.service.RootMeanSquare(req.GetFirstNumber(), req.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) GeometricMean(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := s.service.GeometricMean(req.GetFirstNumber(), req.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) BodyMassIndex(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := s.service.BodyMassIndex(req.GetFirstNumber(), req.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) PowerLevelDiff(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := s.service.PowerLevelDiff(req.GetFirstNumber(), req.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) PercentageValueChange(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := s.service.PercentageValueChange(req.GetFirstNumber(), req.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}
