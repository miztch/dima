package main

type Match struct {
	Name             string `json:"matchName" dynamodbav:"matchName"`
	StartTime        string `json:"startTime" dynamodbav:"startTime"`
	EventName        string `json:"eventName" dynamodbav:"eventName"`
	Teams            []Team `json:"teams" dynamodbav:"teams"`
	StartDate        string `json:"startDate" dynamodbav:"startDate"`
	PagePath         string `json:"pagePath" dynamodbav:"pagePath"`
	BestOf           int    `json:"bestOf" dynamodbav:"bestOf"`
	Id               int    `json:"id" dynamodbav:"id"`
	EventCountryFlag string `json:"eventCountryFlag" dynamodbav:"eventCountryFlag"`
}

type Team struct {
	Title string `json:"title"`
}
