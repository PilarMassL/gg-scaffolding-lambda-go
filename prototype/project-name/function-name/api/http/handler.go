package http

import (
	"repo-name/project-name/function-name/internal/domain"

	"github.com/aws/aws-lambda-go/events"

	"go.uber.org/zap"
)

type HttpHandler struct {
	service *domain.FunctionNameService
}

func NewHttpHandler(service *domain.FunctionNameService) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (h *HttpHandler) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// events.APIGatewayProxyRequest -> Objeto de Dominio ->
	// Invocar al servicio interno de dominio -> events.APIGatewayProxyResponse

	logger, _ := zap.NewProduction()

	// Invocamos al servicio interno de dominio
	resp, err := h.service.GetMyIp()

	if err != nil {
		logger.Error("Fail obtaining My Ip")
		return events.APIGatewayProxyResponse{}, err
	}

	// transformamos al tipo events.APIGatewayProxyResponse
	return events.APIGatewayProxyResponse{
		Body:       resp,
		StatusCode: 200,
	}, nil
}
