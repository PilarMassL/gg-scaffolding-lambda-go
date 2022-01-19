package dynamo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"repo-name/project-name/function-name/internal/port"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"repo-name/project-name/function-name/internal/domain/model"
)

const tableName = "POCExampleTable"


// NewRepository creates and returns a new repository instance
func NewRepository(database *dynamodb.DynamoDB) port.ExampleClientPort {
	return &repository{
		database: database,
		table:    tableName,
	}
}

type repository struct {
	database *dynamodb.DynamoDB
	table    string
}

func (s *repository) Create(resource *model.Example) (err error) {
	log.Logger.Printf("Adding a new example [%s] ", resource.Key)


	item, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      item,
	})

	log.Logger.Printf("Key %s inserted. ", resource.Key)
	return nil
}

func (s *repository) Update(resource *model.Example) (err error) {
	log.Logger.Printf("Adding a new example [%s] ", resource.Key)

	item, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      item,
	})

	log.Logger.Printf("Key %s inserted.", resource.Key)
	return nil
}

func (s *repository) GetByID(key string) (example *model.Example, err error) {
	log.Logger.Print("Getting Example by Key")

	result, err := s.database.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(s.table),
		Key: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(key),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if result.Item == nil {
		return nil, errors.New("example not found")
	}

	example = &model.Example{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &example)
	if err != nil {
		return
	}

	return
}

func (s repository) Delete(key string) (err error) {
	log.Logger.Printf("Deleting an example [%s] ", key)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(key),
			},
		},
		TableName: aws.String(s.table),
	}

	_, err = s.database.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}

	return nil
}