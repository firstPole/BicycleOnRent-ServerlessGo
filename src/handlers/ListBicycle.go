package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var listOfCycles = []struct {
	ID            int     `json:"id"`
	CycleReview   float32 `json:"CycleReview"`
	AvailableDate string  `json:"AvailableTime"`
	RentAmount    float32 `json:"RentAmount"`
	CurrencyType  string  `json:"CurrencyType"`
	ReturnDate    string  `json:"ReturnDate"`
}{
	{
		ID:            1,
		CycleReview:   3.4,
		AvailableDate: time.Now().String(),
		RentAmount:    400,
		CurrencyType:  "Rs",
		ReturnDate:    time.Now().AddDate(0, 0, 15).String(),
	},
	{
		ID:            2,
		CycleReview:   4.0,
		AvailableDate: time.Now().String(),
		RentAmount:    500,
		CurrencyType:  "Rs",
		ReturnDate:    time.Now().AddDate(0, 0, 15).String(),
	},
	{
		ID:            3,
		CycleReview:   4.4,
		AvailableDate: time.Now().String(),
		RentAmount:    550,
		CurrencyType:  "Rs",
		ReturnDate:    time.Now().AddDate(0, 0, 15).String(),
	},
}

func Handler() (events.APIGatewayProxyResponse, error) {
	response, err := json.Marshal(listOfCycles)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-type": "application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
