package repository

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type userDynamo struct {
	client    *dynamodb.Client
	tableName string
}
