package dynamo

import (
	"context"

	config_pkg "github.com/andreparelho/user-api-crud/pkg/config"
	config_aws "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoClient(ctx context.Context, config config_pkg.Config) (*dynamodb.Client, error) {
	cfg, err := config_aws.LoadDefaultConfig(
		ctx,
		config_aws.WithRegion(config.Dynamo.Region),
	)
	if err != nil {
		return nil, err
	}

	opts := []func(*dynamodb.Options){}

	if endpoint := config.Dynamo.Endpoint; endpoint != "" {
		opts = append(opts, func(o *dynamodb.Options) {
			o.BaseEndpoint = &endpoint
		})
	}

	return dynamodb.NewFromConfig(cfg, opts...), nil
}
