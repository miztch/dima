package main

type match struct {
	Name             string `dynamodbabv:"matchName"`
	StartTime        string `dynamodbabv:"startTime"`
	EventName        string `dynamodbabv:"eventName"`
	Teams            []team `dynamodbabv:"teams"`
	StartDate        string `dynamodbabv:"startDate"`
	PagePath         string `dynamodbabv:"pagePath"`
	BestOf           int    `dynamodbabv:"bestOf"`
	Id               int    `dynamodbabv:"id"`
	EventCountryFlag string `dynamodbabv:"eventCountryFlag"`
}

type team struct {
	Title string
}
