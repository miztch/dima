package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDBClient(ctx context.Context, tableName string) (*DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("[error] failed to load configuration, %w", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{client: client, tableName: tableName}, nil
}

func (d *DynamoDBClient) QueryMatchesByStartDate(ctx context.Context, startDate string) (matches []match, err error) {
	keyEx := expression.Key("startDate").Equal(expression.Value(startDate))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, fmt.Errorf("[error] failed to build expression, %w", err)
	}

	queryPaginator := dynamodb.NewQueryPaginator(d.client, &dynamodb.QueryInput{
		TableName:                 &d.tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	for queryPaginator.HasMorePages() {
		response, err := queryPaginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("[error] failed to query matches by start date, %w", err)
		}

		var matchPage []match
		err = attributevalue.UnmarshalListOfMaps(response.Items, &matchPage)
		if err != nil {
			return nil, fmt.Errorf("[error] failed to unmarshal match page, %w", err)
		}

		matches = append(matches, matchPage...)
	}
	return matches, err
}
