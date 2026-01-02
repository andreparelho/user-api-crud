package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"github.com/andreparelho/user-api-crud/pkg/config"
	"github.com/andreparelho/user-api-crud/pkg/dynamo"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserRepository interface {
	Save(ctx context.Context, user dynamo.User) error
	GetByID(ctx context.Context, id string) (*dynamo.User, error)
}

func NewUserRepository(client *dynamodb.Client, cfg config.Config) UserRepository {
	return &userDynamo{
		client:    client,
		tableName: cfg.Dynamo.TableName,
	}
}

func (u *userDynamo) Save(ctx context.Context, user dynamo.User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	_, err = u.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &u.tableName,
		Item:      item,
	})
	return err
}

func (u *userDynamo) GetByID(ctx context.Context, id string) (*dynamo.User, error) {
	out, err := u.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &u.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}

	if out.Item == nil {
		return nil, nil
	}

	var user dynamo.User
	if err := attributevalue.UnmarshalMap(out.Item, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
