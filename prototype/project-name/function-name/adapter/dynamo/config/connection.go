package config

import "log"


import (
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoDataBase struct {
	databaseConnection *dynamodb.DynamoDB
}

// NewDynamoDBStorage creates and returns a new Lock dynamo db connection instance
func NewDynamoDBStorage() (DataBase, error) {
	log.Logger.Print("New instance Dynamo storage")
	dataBase := &dynamoDataBase{}
	err := dataBase.OpenConnection()
	if err != nil {
		return nil, err
	}
	return dataBase, nil
}

// OpenConnection start dynamo db connection
func (db *dynamoDataBase) OpenConnection() error {
	log.Logger.Print("Starting Dynamo connection")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	db.databaseConnection = dynamodb.New(sess)
	log.Logger.Print("DynamoDB UP")
	return nil
}

// GetConnection get dynamo db connection
func (db *dynamoDataBase) GetConnection() interface{} {
	return db.databaseConnection
}
