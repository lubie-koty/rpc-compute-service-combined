package http

import (
	"net/http"

	"github.com/lubie-koty/rpc-compute-service-combined/internal/core/types"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/util"
)

type HTTPService struct {
	service types.MathService
}

func NewHTTPService(service types.MathService) *HTTPService {
	return &HTTPService{
		service: service,
	}
}

type UnaryRequest struct {
	Number float64 `json:"number" validate:"required"`
}

type BinaryRequest struct {
	FirstNumber  float64 `json:"first_number" validate:"required"`
	SecondNumber float64 `json:"second_number" validate:"required"`
}

type OperationResponse struct {
	Result float64 `json:"result" validate:"required"`
}

func (s *HTTPService) Sqrt(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[UnaryRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.Sqrt(body.Number)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) Abs(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[UnaryRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.Abs(body.Number)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) Power(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[BinaryRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.Power(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) Log(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[BinaryRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.Log(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) Round(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[BinaryRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.Round(body.FirstNumber, int64(body.SecondNumber))
	util.WriteResponse(w, OperationResponse{Result: result})
}
