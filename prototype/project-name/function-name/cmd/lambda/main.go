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
	logger.Info("Iniciando componentes")
	//Aquí se hace la inyección de las dependencias.

	//Se crean los adaptadores
	myIpClientUsingAws := adapter.NewMyIpClientUsingAws()
	myStorageUsingS3 := adapter.NewMyStorageUsingS3()

	//El servicio raíz de la función
	myFunctionNameService := domain.NewFunctionNameService(myIpClientUsingAws, myStorageUsingS3)

	// El handler que manejará el evento que dispara la función. Ej: evento s3, apigw trigger.
	myHandler := s3.NewS3EventHandler(myFunctionNameService)
	//myHandler := http.NewHttpHandler(myFunctionNameService)

	logger.Info("Iniciando función")
	lambda.Start(myHandler.Handler)
}
