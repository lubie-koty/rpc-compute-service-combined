package grpc

import (
	"context"
	"io"
	"log/slog"

	"github.com/lubie-koty/rpc-compute-service-combined/internal/core/types"
	pb "github.com/lubie-koty/rpc-compute-service-combined/protos"
	"google.golang.org/grpc"
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

func (s *GRPCService) RootMeanSquare(stream grpc.ClientStreamingServer[pb.RepeatedOperationRequest, pb.OperationResponse]) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			result, err := s.service.RootMeanSquare(request.GetNumbers())
			if err != nil {
				s.logger.Error(err.Error())
				return status.Error(codes.Internal, err.Error())
			}
			return stream.SendAndClose(&pb.OperationResponse{Result: result})
		}
		if err != nil {
			s.logger.Error(err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}
}

func (s *GRPCService) GeometricMean(req *pb.RepeatedOperationRequest, stream grpc.ServerStreamingServer[pb.OperationResponse]) error {
	result, err := s.service.GeometricMean(req.GetNumbers())
	if err != nil {
		s.logger.Error(err.Error())
		return status.Error(codes.Internal, err.Error())
	}
	if err := stream.Send(&pb.OperationResponse{Result: result}); err != nil {
		s.logger.Error(err.Error())
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (s *GRPCService) BodyMassIndex(stream grpc.BidiStreamingServer[pb.OperationRequest, pb.OperationResponse]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			s.logger.Error(err.Error())
			return status.Error(codes.Internal, err.Error())
		}
		result, err := s.service.BodyMassIndex(req.GetFirstNumber(), req.GetSecondNumber())
		if err != nil {
			s.logger.Error(err.Error())
			return status.Error(codes.Internal, err.Error())
		}
		if err := stream.Send(&pb.OperationResponse{Result: result}); err != nil {
			s.logger.Error(err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}
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
