package main

import "os"

type dynamoDBConfig struct {
	TableName   string
	Region      *string
	EndpointURL *string
}

func getDynamoDBConfig() dynamoDBConfig {
	region := os.Getenv("AWS_REGION")
	endpointURL := os.Getenv("DYNAMODB_ENDPOINT_URL")
	return dynamoDBConfig{
		TableName:   os.Getenv("TABLE_NAME"),
		Region:      &region,
		EndpointURL: &endpointURL,
	}
}
