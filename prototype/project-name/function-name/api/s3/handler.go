package s3

import (
	"context"
	"errors"
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

func (h *S3EventHandler) Handler(ctx context.Context, event events.S3Event) error {
	// events.S3Event -> Objeto de Dominio ->
	// Invocar al servicio interno de dominio

	logger, _ := zap.NewProduction()

	//Obtenemos el nombre del archivo que se agregó:
	newFile, err := extractObjectKey(event)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	// Invocamos al servicio interno de dominio
	errCipher := h.service.AppendTimestamp(newFile)

	if errCipher != nil {
		logger.Error(errCipher.Error())
		return errCipher
	}
	logger.Info("Se terminó el proceso correctamente")
	return nil
}

func extractObjectKey(event events.S3Event) (string, error) {
	records := event.Records
	for _, r := range records {
		return r.S3.Object.Key, nil //Enviamos el primero que encontramos
	}
	return "", errors.New("no key found")
}
