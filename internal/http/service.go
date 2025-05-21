package http

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/core/types"
	"github.com/lubie-koty/rpc-compute-service-combined/internal/util"
)

type HTTPService struct {
	service types.MathService
	logger  *slog.Logger
}

func NewHTTPService(service types.MathService, logger *slog.Logger) *HTTPService {
	return &HTTPService{
		service: service,
		logger:  logger,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type UnaryRequest struct {
	Number float64 `json:"number" validate:"required"`
}

type BinaryRequest struct {
	FirstNumber  float64 `json:"first_number" validate:"required"`
	SecondNumber float64 `json:"second_number" validate:"required"`
}

type RepeatedOperationRequest struct {
	Numbers []float64 `json:"numbers" validate:"required"`
}

type OperationRequest struct {
	FirstNumber  float64 `json:"first_number" validate:"required"`
	SecondNumber float64 `json:"second_number" validate:"required"`
}

type OperationResponse struct {
	Result float64 `json:"result" validate:"required"`
}

func (s *HTTPService) RootMeanSquare(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "an error occurred while upgrading connection", http.StatusInternalServerError)
		s.logger.Error(err.Error())
		return
	}
	defer conn.Close()

	body, err := util.GetWebsocketResponse[RepeatedOperationRequest](conn)
	if err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}

	result, err := s.service.RootMeanSquare(body.Numbers)
	if err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}

	if err := util.WriteWebsocketResponse(conn, OperationResponse{Result: result}); err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}
}

func (s *HTTPService) GeometricMean(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "an error occurred while upgrading connection", http.StatusInternalServerError)
		s.logger.Error(err.Error())
		return
	}
	defer conn.Close()

	body, err := util.GetWebsocketResponse[RepeatedOperationRequest](conn)
	if err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}

	result, err := s.service.GeometricMean(body.Numbers)
	if err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}

	if err := util.WriteWebsocketResponse(conn, OperationResponse{Result: result}); err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}
}

func (s *HTTPService) BodyMassIndex(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "an error occurred while upgrading connection", http.StatusInternalServerError)
		s.logger.Error(err.Error())
		return
	}
	defer conn.Close()

	body, err := util.GetWebsocketResponse[OperationRequest](conn)
	if err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}

	result, err := s.service.BodyMassIndex(body.FirstNumber, body.SecondNumber)
	if err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}

	if err := util.WriteWebsocketResponse(conn, OperationResponse{Result: result}); err != nil {
		util.WriteWebsocketErrorResponse(conn, err)
		s.logger.Error(err.Error())
		return
	}
}

func (s *HTTPService) PowerLevelDiff(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result, err := s.service.PowerLevelDiff(body.FirstNumber, body.SecondNumber)
	if err != nil {
		http.Error(w, "an error occurred while computing result", http.StatusInternalServerError)
		s.logger.Error(err.Error())
		return
	}
	util.WriteResponse(w, OperationResponse{Result: result})
}

func (s *HTTPService) PercentageValueChange(w http.ResponseWriter, r *http.Request) {
	util.ValidateRequest(w, r)
	body, err := util.GetRequestBody[OperationRequest](w, r)
	if err != nil {
		return
	}
	result, err := s.service.PercentageValueChange(body.FirstNumber, body.SecondNumber)
	if err != nil {
		http.Error(w, "an error occurred while computing result", http.StatusInternalServerError)
		s.logger.Error(err.Error())
		return
	}
	util.WriteResponse(w, OperationResponse{Result: result})
}
