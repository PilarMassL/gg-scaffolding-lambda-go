package s3

import (
	"context"
	"encoding/json"
	"errors"
	"repo-name/project-name/function-name/internal/domain"

	"github.com/aws/aws-lambda-go/events"

	"go.uber.org/zap"
)

type eventBridgeHandler struct {
	service *domain.FunctionNameService
}

type s3EventBridgeData struct {
	Version string `json:"version"`
	Bucket  struct {
		Name string `json:"name"`
	} `json:"bucket"`
	Object struct {
		Key        string `json:"key"`
		Size       int32  `json:"size"`
		Etag       string `json:"etag"`
		Version_id string `json:"version-id"`
	} `json:"object"`
	Request_id              string `json:"request-id"`
	Requester               string `json:"requester"`
	Destination_access_tier string `json:"destination-access-tier"`
}

func NewS3EventHandler(service *domain.FunctionNameService) *eventBridgeHandler {
	return &eventBridgeHandler{
		service: service,
	}
}

func (h *eventBridgeHandler) Handler(ctx context.Context, event events.CloudWatchEvent) error {
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

func extractObjectKey(event events.CloudWatchEvent) (string, error) {

	var detailsS3 *s3EventBridgeData
	if event.Source == "aws.s3" {
		event.Detail.MarshalJSON()
		err := json.Unmarshal(event.Detail , detailsS3)
		if err == nil {
			return detailsS3.Object.Key, nil
		}
		return "", err
	}
	return "", errors.New("no s3 event received")
}
