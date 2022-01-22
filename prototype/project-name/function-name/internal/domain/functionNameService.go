package domain

import (
	"fmt"
	"log"
	"repo-name/project-name/function-name/internal/port"
	"time"
)

type FunctionNameService struct {
	myIpClient port.MyIpClientPort
	myStorage  port.MyStoragePort
}

func NewFunctionNameService(client port.MyIpClientPort, storage port.MyStoragePort) *FunctionNameService {
	return &FunctionNameService{
		myIpClient: client,
		myStorage:  storage,
	}
}

func (s *FunctionNameService) GetMyIp() (string, error) {
	return s.myIpClient.GetIp()
}

func (s *FunctionNameService) AppendTimestamp(filename string) error {
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
