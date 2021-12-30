package main

import (
	"repo-name/project-name/function-name/adapter"
	"repo-name/project-name/function-name/api/http"
	"repo-name/project-name/function-name/internal/domain"

	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("Lambda starts")
	//Aquí se hace la inyección de las dependencias.
	myIpClientUsingAws := adapter.NewMyIpClientUsingAws()
	myFunctionNameService := domain.NewFunctionNameService(myIpClientUsingAws)
	myHandler := http.NewHttpHandler(myFunctionNameService)

	lambda.Start(myHandler.Handler)
}
