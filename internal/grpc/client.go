package grpc

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbComplex "github.com/lubie-koty/rpc-compute-service-combined/protos/complex"
	pbSimple "github.com/lubie-koty/rpc-compute-service-combined/protos/simple"
)

type GRPCSimpleServiceClient struct {
	ctx    context.Context
	logger *slog.Logger
	client pbSimple.SimpleComputeClient
}

type GRPCComplexServiceClient struct {
	ctx    context.Context
	logger *slog.Logger
	client pbComplex.ComplexComputeClient
}

func NewGRPCSimpleServiceClient(ctx context.Context, logger *slog.Logger, address string) (*GRPCSimpleServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pbSimple.NewSimpleComputeClient(conn)
	return &GRPCSimpleServiceClient{
		ctx:    ctx,
		logger: logger,
		client: client,
	}, conn
}

func NewGRPCComplexServiceClient(ctx context.Context, logger *slog.Logger, address string) (*GRPCComplexServiceClient, *grpc.ClientConn) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pbComplex.NewComplexComputeClient(conn)
	return &GRPCComplexServiceClient{
		ctx:    ctx,
		logger: logger,
		client: client,
	}, conn

}

func (s *GRPCSimpleServiceClient) Add(numbers []float64) (float64, error) {
	result, err := s.client.Add(s.ctx, &pbSimple.OperationRequest{Numbers: numbers})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}

func (s *GRPCSimpleServiceClient) Sub(numbers []float64) (float64, error) {
	result, err := s.client.Sub(s.ctx, &pbSimple.OperationRequest{Numbers: numbers})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}

func (s *GRPCSimpleServiceClient) Mul(numbers []float64) (float64, error) {
	result, err := s.client.Mul(s.ctx, &pbSimple.OperationRequest{Numbers: numbers})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}

func (s *GRPCSimpleServiceClient) Div(numbers []float64) (float64, error) {
	result, err := s.client.Div(s.ctx, &pbSimple.OperationRequest{Numbers: numbers})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}

func (c *GRPCComplexServiceClient) Sqrt(num float64) (float64, error) {
	result, err := c.client.Sqrt(c.ctx, &pbComplex.UnaryRequest{Number: num})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}

func (c *GRPCComplexServiceClient) Abs(num float64) (float64, error) {
	result, err := c.client.Abs(c.ctx, &pbComplex.UnaryRequest{Number: num})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}

func (c *GRPCComplexServiceClient) Power(base float64, exponent float64) (float64, error) {
	result, err := c.client.Power(c.ctx, &pbComplex.BinaryRequest{FirstNumber: base, SecondNumber: exponent})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}

func (c *GRPCComplexServiceClient) Log(num float64, base float64) (float64, error) {
	result, err := c.client.Log(c.ctx, &pbComplex.BinaryRequest{FirstNumber: num, SecondNumber: base})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}

func (c *GRPCComplexServiceClient) Round(num float64, precision float64) (float64, error) {
	result, err := c.client.Round(c.ctx, &pbComplex.BinaryRequest{FirstNumber: num, SecondNumber: precision})
	if err != nil {
		return 0, err
	}
	return result.Result, nil
}
