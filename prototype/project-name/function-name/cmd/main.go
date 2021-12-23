package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"repo-name/project-name/function-name/api/http"
)

func main() {
	fmt.Print("Hola mundo desde main")
	lambda.Start(http.Handler)
}
