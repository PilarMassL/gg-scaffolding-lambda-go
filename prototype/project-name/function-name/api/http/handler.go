package http

import (
	"github.com/aws/aws-lambda-go/events"

	"repo-name/project-name/function-name/internal/domain"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// events.APIGatewayProxyRequest -> Objeto de Dominio ->
	// Invocar al servicio interno de dominio -> events.APIGatewayProxyResponse

	resp, err := domain.GetIp()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       resp,
		StatusCode: 200,
	}, nil
}
