package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Dynamo struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type DynamoDBInterface interface {
	ScanAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error)
	GetOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error)
	Query(condition expression.Expression, indexName string, tableName string) (response *dynamodb.QueryOutput, err error)
	CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error)
	Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error)
}

func Init(awsSession *session.Session) DynamoDBInterface {
	con := dynamodb.New(awsSession)
	return &Dynamo{
		connection: con,
		logMode:    false,
	}
}

func (db *Dynamo) ScanAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	}
	return db.connection.Scan(input)
}

func (db *Dynamo) GetOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	}
	return db.connection.GetItem(input)
}

func (db *Dynamo) Query(condition expression.Expression, indexName string, tableName string) (response *dynamodb.QueryOutput, err error) {
	query := &dynamodb.QueryInput{
		IndexName:                 aws.String(indexName),
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    condition.KeyCondition(),
		ExpressionAttributeValues: condition.Values(),
		ExpressionAttributeNames:  condition.Names(),
	}
	return db.connection.Query(query)
}

func (db *Dynamo) CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db *Dynamo) Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.DeleteItem(input)
}
