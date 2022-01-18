package main

import (
	"repo-name/project-name/function-name/adapter"
	"repo-name/project-name/function-name/api/s3"
	"repo-name/project-name/function-name/internal/domain"

	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("Lambda starts")
	//Aquí se hace la inyección de las dependencias.
	//myIpClientUsingAws := adapter.NewMyIpClientUsingAws()
	//myFunctionNameService := domain.NewFunctionNameService(myIpClientUsingAws)
	//myHttpHandler := http.NewHttpHandler(myFunctionNameService)
	//lambda.Start(myHttpHandler.Handler)

	myStorageUsingS3 := adapter.NewMyStorageUsingS3()
	myFunctionNameDemoS3Service := domain.NewFunctionNameDemoS3Service(myStorageUsingS3)
	myS3EventHandler := s3.NewS3EventHandler(myFunctionNameDemoS3Service)

	lambda.Start(myS3EventHandler.Handler)
}
