package s3

import (
	"log"
	"repo-name/project-name/function-name/internal/domain"

	"github.com/aws/aws-lambda-go/events"

	"go.uber.org/zap"
)

type S3EventHandler struct {
	service *domain.FunctionNameService
}

func NewS3EventHandler(service *domain.FunctionNameService) *S3EventHandler {
	return &S3EventHandler{
		service: service,
	}
}

func (h *S3EventHandler) Handler(request events.S3Event) error {
	// events.S3Event -> Objeto de Dominio ->
	// Invocar al servicio interno de dominio -> ?

	logger, _ := zap.NewProduction()

	// Invocamos al servicio interno de dominio
	resp, err := h.service.GetMyIp()

	if err != nil {
		logger.Error("Fail obtaining My Ip")
		return err // ?
	}

	// transformamos al tipo events.?
	log.Printf("La respuesta obtenida del servicio es: %s", resp)
	return nil
}
