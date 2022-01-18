package domain

import (
	"fmt"
	"log"
	"repo-name/project-name/function-name/internal/port"
	"time"
)

type FunctionNameDemoS3Service struct {
	myStorage port.MyStoragePort
}

func NewFunctionNameDemoS3Service(storage port.MyStoragePort) *FunctionNameDemoS3Service {
	return &FunctionNameDemoS3Service{
		myStorage: storage,
	}
}

func (s *FunctionNameDemoS3Service) AppendTimestamp(filename string) error {
	log.Printf("Intentamos obtener el archivo: %s", filename)
	// Intentamos abrir el archivo
	content, err := s.myStorage.GetFile(filename, "1")
	if err != nil {
		return err
	}
	// Creamos mensaje de cola con el timestamp
	tail := []byte(fmt.Sprintf("Recibido: %s", time.Now().Format(time.RFC850)))
	finalContent := appendTail(content, tail)

	// Intentamos guardar el archivo con el timestamp
	errPuttingFile := s.myStorage.PutFile(fmt.Sprintf("received/%s", filename), finalContent)
	if errPuttingFile != nil {
		return err
	}

	return nil
}

func appendTail(content []byte, tail []byte) []byte {
	result := make([]byte, len(content)+len(tail))
	result = append(result, content...)
	result = append(result, tail...)
	return result
}
