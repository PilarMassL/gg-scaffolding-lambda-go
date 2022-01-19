package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
	"repo-name/project-name/function-name/api/http/dynamo/create_example"
	"repo-name/project-name/function-name/cmd/lambda/dynamo"
)

func main() {
	logger, _ := zap.NewProduction()
	logger.Info("Lambda Dynamo Create starts")
	myHandler := create_example.NewHandler(dynamo.Service)
	lambda.Start(myHandler.Handler)
}

