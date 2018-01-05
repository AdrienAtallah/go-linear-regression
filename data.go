package main

type PriceRequest struct {
	CurrentDate    int    `json:"currentDate"`
	PredictionDate string `json:"predictionDate"`
	Currency       string `json:"currency"`
}
