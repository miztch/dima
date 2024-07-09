package main

type match struct {
	Name             string `json:"matchName" dynamodbabv:"matchName"`
	StartTime        string `json:"startTime" dynamodbabv:"startTime"`
	EventName        string `json:"eventName" dynamodbabv:"eventName"`
	Teams            []team `json:"teams" dynamodbabv:"teams"`
	StartDate        string `json:"startDate" dynamodbabv:"startDate"`
	PagePath         string `json:"pagePath" dynamodbabv:"pagePath"`
	BestOf           int    `json:"bestOf" dynamodbabv:"bestOf"`
	Id               int    `json:"id" dynamodbabv:"id"`
	EventCountryFlag string `json:"eventCountryFlag" dynamodbabv:"eventCountryFlag"`
}

type team struct {
	Title string `json:"title"`
}
