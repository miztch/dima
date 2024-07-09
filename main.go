package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

var app *App
var gorillaMuxAdapter *gorillamux.GorillaMuxAdapter

func init() {
	log.Printf("Cold start")
	app = NewApp()
	gorillaMuxAdapter = gorillamux.New(app.Router.muxRouter)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r, err := gorillaMuxAdapter.ProxyWithContext(ctx, *core.NewSwitchableAPIGatewayRequestV1(&req))
	return *r.Version1(), err
}

func main() {
	lambda.Start(Handler)
}
