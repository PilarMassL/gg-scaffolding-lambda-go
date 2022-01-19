package create_example

import (
	"encoding/json"
	"log"
	"repo-name/project-name/function-name/internal/domain"
	"repo-name/project-name/function-name/internal/domain/model"

	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	service domain.ExampleCRUDService
}

func NewHandler(service domain.ExampleCRUDService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger.Printf("Calling Lambda Create Example %s",  request.Body)

	data := model.Example{}
	err:= json.Unmarshal([]byte(request.Body), &data)
	if err != nil {
		log.Logger.Print("Fail unmarshaling request body")
		return events.APIGatewayProxyResponse{}, err
	}
	err= h.service.Create(&data)
	if err != nil {
		log.Logger.Print("Fail creating example")
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil


}
