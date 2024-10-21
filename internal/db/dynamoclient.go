package db

import (
	"ekoa-certificate-generator/config"
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type DynamoDB struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type IDynamoDB interface {
	ScanAll(condition expression.Expression, tableName string, lastEvaluatedKey map[string]*dynamodb.AttributeValue) (response *dynamodb.ScanOutput, err error)
	GetOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error)
	Query(condition expression.Expression, indexName string, tableName string) (response *dynamodb.QueryOutput, err error)
	CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error)
	Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error)
	ToString(input map[string]*dynamodb.AttributeValue) (string, error)
	ToAttributeValue(input string) (map[string]*dynamodb.AttributeValue, error)
}

func NewClient(c config.AWS) (IDynamoDB, error) {
	sess, err := config.CreateAWSSession(c)
	if err != nil {
		return nil, err
	}

	con := dynamodb.New(sess)
	return &DynamoDB{
		connection: con,
		logMode:    false,
	}, nil
}

func (db *DynamoDB) ScanAll(condition expression.Expression, tableName string, lastEvaluatedKey map[string]*dynamodb.AttributeValue) (response *dynamodb.ScanOutput, err error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		ExclusiveStartKey:         lastEvaluatedKey,
		TableName:                 aws.String(tableName),
	}

	return db.connection.Scan(input)
}

func (db *DynamoDB) GetOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {
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

func (db *DynamoDB) Query(condition expression.Expression, indexName string, tableName string) (response *dynamodb.QueryOutput, err error) {
	query := &dynamodb.QueryInput{
		IndexName:                 aws.String(indexName),
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    condition.KeyCondition(),
		ExpressionAttributeValues: condition.Values(),
		ExpressionAttributeNames:  condition.Names(),
	}
	return db.connection.Query(query)
}

func (db *DynamoDB) CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {
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

func (db *DynamoDB) Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {
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

func (db *DynamoDB) ToString(input map[string]*dynamodb.AttributeValue) (string, error) {
	var inputMap map[string]interface{}
	err := dynamodbattribute.UnmarshalMap(input, &inputMap)
	if err != nil {
		return "", err
	}
	bytesJSON, err := json.Marshal(inputMap)
	if err != nil {
		return "", err
	}
	output := base64.StdEncoding.EncodeToString(bytesJSON)
	return output, nil
}

func (db *DynamoDB) ToAttributeValue(input string) (map[string]*dynamodb.AttributeValue, error) {
	bytesJSON, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	outputJSON := map[string]interface{}{}
	err = json.Unmarshal(bytesJSON, &outputJSON)
	if err != nil {
		return nil, err
	}

	return dynamodbattribute.MarshalMap(outputJSON)
}
