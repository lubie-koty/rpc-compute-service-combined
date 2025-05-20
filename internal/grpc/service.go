package grpc

import (
	"context"
	"log/slog"

	"github.com/lubie-koty/rpc-compute-service-combined/internal/core/types"
	pb "github.com/lubie-koty/rpc-compute-service-combined/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCService struct {
	pb.UnimplementedCombinedComputeServer
	service types.MathService
	logger  *slog.Logger
}

func NewGRPCService(service types.MathService, logger *slog.Logger) *GRPCService {
	return &GRPCService{
		service: service,
		logger:  logger,
	}
}

func (s *GRPCService) RootMeanSquare(ctx context.Context, req *pb.RepeatedOperationRequest) (*pb.OperationResponse, error) {
	result, err := s.service.RootMeanSquare(req.GetNumbers())
	if err != nil {
		s.logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) GeometricMean(ctx context.Context, req *pb.RepeatedOperationRequest) (*pb.OperationResponse, error) {
	result, err := s.service.GeometricMean(req.GetNumbers())
	if err != nil {
		s.logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) BodyMassIndex(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result, err := s.service.BodyMassIndex(req.GetFirstNumber(), req.GetSecondNumber())
	if err != nil {
		s.logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) PowerLevelDiff(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result, err := s.service.PowerLevelDiff(req.GetFirstNumber(), req.GetSecondNumber())
	if err != nil {
		s.logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.OperationResponse{Result: result}, nil
}

func (s *GRPCService) PercentageValueChange(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	result, err := s.service.PercentageValueChange(req.GetFirstNumber(), req.GetSecondNumber())
	if err != nil {
		s.logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.OperationResponse{Result: result}, nil
}
