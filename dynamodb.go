package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
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

func (d *DynamoDBClient) QueryMatchesByStartDate(ctx context.Context, startDate string) (matches []Match, err error) {
	keyEx := expression.Key("startDate").Equal(expression.Value(startDate))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, fmt.Errorf("[error] failed to build expression, %w", err)
	}

	queryPaginator := dynamodb.NewQueryPaginator(d.client, &dynamodb.QueryInput{
		TableName:                 aws.String(d.tableName),
		IndexName:                 aws.String("startDate-id-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	matches = []Match{} // Initialize matches to an empty slice to avoid nil pointer dereference

	for queryPaginator.HasMorePages() {
		response, err := queryPaginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("[error] failed to query matches by start date, %w", err)
		}

		var matchPage []Match
		err = attributevalue.UnmarshalListOfMaps(response.Items, &matchPage)
		if err != nil {
			return nil, fmt.Errorf("[error] failed to unmarshal match page, %w", err)
		}

		matches = append(matches, matchPage...)
	}

	return matches, nil
}
