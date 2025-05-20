package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/lubie-koty/rpc-compute-service-combined/internal/util"
)

type HTTPSimpleServiceClient struct {
	ctx     *context.Context
	logger  *slog.Logger
	client  *http.Client
	address string
}

type HTTPComplexServiceClient struct {
	ctx     *context.Context
	logger  *slog.Logger
	client  *http.Client
	address string
}

func NewHTTPSimpleServiceClient(ctx *context.Context, logger *slog.Logger, address string) *HTTPSimpleServiceClient {
	return &HTTPSimpleServiceClient{
		ctx:     ctx,
		logger:  logger,
		client:  &http.Client{},
		address: address,
	}
}

func NewHTTPComplexServiceClient(ctx *context.Context, logger *slog.Logger, address string) *HTTPComplexServiceClient {
	return &HTTPComplexServiceClient{
		ctx:     ctx,
		logger:  logger,
		client:  &http.Client{},
		address: address,
	}
}

func (s *HTTPSimpleServiceClient) Add(numbers []float64) (float64, error) {
	req, err := util.CreateRequest[RepeatedOperationRequest](
		"POST",
		fmt.Sprintf("%s/add", s.address),
		RepeatedOperationRequest{Numbers: numbers},
	)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil
}

func (s *HTTPSimpleServiceClient) Sub(numbers []float64) (float64, error) {
	req, err := util.CreateRequest[RepeatedOperationRequest](
		"POST",
		fmt.Sprintf("%s/sub", s.address),
		RepeatedOperationRequest{Numbers: numbers},
	)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil
}

func (s *HTTPSimpleServiceClient) Mul(numbers []float64) (float64, error) {
	req, err := util.CreateRequest[RepeatedOperationRequest](
		"POST",
		fmt.Sprintf("%s/mul", s.address),
		RepeatedOperationRequest{Numbers: numbers},
	)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil
}

func (s *HTTPSimpleServiceClient) Div(numbers []float64) (float64, error) {
	req, err := util.CreateRequest[RepeatedOperationRequest](
		"POST",
		fmt.Sprintf("%s/div", s.address),
		RepeatedOperationRequest{Numbers: numbers},
	)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil
}

func (c *HTTPComplexServiceClient) Sqrt(num float64) (float64, error) {
	req, err := util.CreateRequest[UnaryRequest](
		"POST",
		fmt.Sprintf("%s/sqrt", c.address),
		UnaryRequest{Number: num},
	)
	if err != nil {
		return 0, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil
}

func (c *HTTPComplexServiceClient) Abs(num float64) (float64, error) {
	req, err := util.CreateRequest[UnaryRequest](
		"POST",
		fmt.Sprintf("%s/abs", c.address),
		UnaryRequest{Number: num},
	)
	if err != nil {
		return 0, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil

}

func (c *HTTPComplexServiceClient) Power(base float64, exponent float64) (float64, error) {
	req, err := util.CreateRequest[BinaryRequest](
		"POST",
		fmt.Sprintf("%s/power", c.address),
		BinaryRequest{FirstNumber: base, SecondNumber: exponent},
	)
	if err != nil {
		return 0, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil
}

func (c *HTTPComplexServiceClient) Log(num float64, base float64) (float64, error) {
	req, err := util.CreateRequest[BinaryRequest](
		"POST",
		fmt.Sprintf("%s/power", c.address),
		BinaryRequest{FirstNumber: num, SecondNumber: base},
	)
	if err != nil {
		return 0, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil
}

func (c *HTTPComplexServiceClient) Round(num float64, precision float64) (float64, error) {
	req, err := util.CreateRequest[BinaryRequest](
		"POST",
		fmt.Sprintf("%s/power", c.address),
		BinaryRequest{FirstNumber: num, SecondNumber: precision},
	)
	if err != nil {
		return 0, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	data, err := util.GetResponseBody[OperationResponse](resp)
	if err != nil {
		return 0, err
	}
	return data.Result, nil
}
