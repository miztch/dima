package main

import (
	"context"
	"log"
)

type App struct {
	Router   *Router
	DBClient *DynamoDBClient
	Handlers *Handlers
}

func NewApp() *App {
	app := &App{}
	app.initDBClient()
	app.initHandlers()
	app.initRouter()
	return app
}

func (a *App) initDBClient() {
	var err error
	a.DBClient, err = NewDynamoDBClient(context.Background(), getDynamoDBConfig().TableName)
	if err != nil {
		log.Printf("[error] failed to create DynamoDB client, %v", err)
	}
}

func (a *App) initHandlers() {
	a.Handlers = NewHandlers(a.DBClient)
}

func (a *App) initRouter() {
	a.Router = NewRouter(a.Handlers)
}
