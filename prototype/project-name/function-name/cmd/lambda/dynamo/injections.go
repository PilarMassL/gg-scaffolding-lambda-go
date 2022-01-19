package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"

	"repo-name/project-name/function-name/adapter/dynamo"
	"repo-name/project-name/function-name/adapter/dynamo/config"
	"repo-name/project-name/function-name/internal/domain"
)

var Service = MakeDependencyInjection()

// MakeDependencyInjection Initialize all dependencies
func MakeDependencyInjection() domain.ExampleCRUDService {
	log.Logger.Print("Start bootstrap app objects")
	database, err := config.NewDynamoDBStorage()
	if err != nil {
		panic(err)
	}
	repository := dynamo.NewRepository(database.GetConnection().(*dynamodb.DynamoDB))
	service := domain.NewExampleCRUDService(repository)
	log.Logger.Print("End bootstrap app objects")
	return service
}
