package main

import (
	"repo-name/project-name/function-name/api/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(http.Handler)
}
