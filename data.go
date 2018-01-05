package main

import "time"

type PriceRequest struct {
	CurrentDate    time.Time `json:"currentDate"`
	PredictionDate string    `json:"predictionDate"`
	Currency       string    `json:"currency"`
}

type PriceResponse struct {
	CurrentDate    time.Time `json:"currentDate"`
	PredictionDate time.Time `json:"predictionDate"`
	Currency       string    `json:"currency"`
	Value          int       `json:"value"`
}
