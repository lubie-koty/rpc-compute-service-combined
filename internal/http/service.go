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

type OperationRequest struct {
	FirstNumber  float64 `json:"first_number" validate:"required"`
	SecondNumber float64 `json:"second_number" validate:"required"`
}

type OperationResponse struct {
	Result float64 `json:"result" validate:"required"`
}

func (s *HTTPService) RootMeanSquare(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.RootMeanSquare(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) GeometricMean(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.GeometricMean(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) BodyMassIndex(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.BodyMassIndex(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) PowerLevelDiff(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.PowerLevelDiff(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) PercentageValueChange(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result := s.service.PercentageValueChange(body.FirstNumber, body.SecondNumber)
	util.WriteResponse(w, OperationResponse{Result: result})
}
