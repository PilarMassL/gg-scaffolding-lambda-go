package domain

import (
	"repo-name/project-name/function-name/internal/port"
)

type FunctionNameService struct {
	myIpClient port.MyIpClientPort
}

func NewFunctionNameService(client port.MyIpClientPort) *FunctionNameService {
	return &FunctionNameService{
		myIpClient: client,
	}
}

func (s *FunctionNameService) GetMyIp() (string, error) {
	return s.myIpClient.GetIp()
}
