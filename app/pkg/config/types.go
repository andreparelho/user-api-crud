package config

type Config struct {
	AppName string
	Port    string
	Env     string
	Dynamo  Dynamo
}

type Dynamo struct {
	Endpoint  string
	Region    string
	TableName string
}
