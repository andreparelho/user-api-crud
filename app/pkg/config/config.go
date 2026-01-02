package config

import (
	"fmt"
	"os"
)

func Load() (*Config, error) {
	var appName, port, env, dynamoEndpoint, dynamoTable, awsRegion string

	appName, errApp := getEnv("APP_NAME")
	if errApp != nil {
		return nil, errApp
	}

	port, errPort := getEnv("PORT")
	if errPort != nil {
		return nil, errPort
	}

	env, errEnv := getEnv("ENV")
	if errEnv != nil {
		return nil, errEnv
	}

	dynamoEndpoint, errDynamoEndpoint := getEnv("DYNAMO_ENDPOINT")
	if errDynamoEndpoint != nil {
		return nil, errDynamoEndpoint
	}

	dynamoTable, errDynamoTable := getEnv("DYNAMO_TABLE_NAME")
	if errDynamoTable != nil {
		return nil, errDynamoTable
	}

	awsRegion, errRegion := getEnv("AWS_REGION")
	if errRegion != nil {
		return nil, errRegion
	}

	return &Config{
		AppName: appName,
		Port:    port,
		Env:     env,
		Dynamo: Dynamo{
			Endpoint:  dynamoEndpoint,
			Region:    awsRegion,
			TableName: dynamoTable,
		},
	}, nil
}

func getEnv(key string) (string, error) {
	if v := os.Getenv(key); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("key: %s is not present in the environment variables file", key)
}
