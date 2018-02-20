package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Welcome!\n")

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	fmt.Println("Go launched at \n", t.Local())

	now := time.Now()
	fmt.Println("The time is now \n", now.Local())

	priceResponse := new(PriceResponse)

	priceResponse.CurrentDate = t
	priceResponse.PredictionDate = now
	priceResponse.Currency = "BTC"
	priceResponse.Value = 12390

	out, err := json.Marshal(priceResponse)

	if err != nil {
		fmt.Println("error marshalling json: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func Plot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// bitStampUsd-full.csv
	// bitStampUsd-2017.csv
	// bitStampUsd-Oct2017.csv
	// bitStampUsd-Dec2017.csv

	filename := "bitStampUsd-Dec2017.csv"
	CsvPlot(filename)

	fmt.Fprintf(w, "plots generated from %s..\n", filename)
}

func Split(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// bitStampUsd-full.csv
	// bitStampUsd-2017.csv
	// bitStampUsd-Oct2017.csv
	// bitStampUsd-Dec2017.csv

	filename := "bitStampUsd-Dec2017.csv"
	CsvSplit(filename)

	fmt.Fprintf(w, "training/testing data sets generated from %s..\n", filename)
}

func Train(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	TrainModel()
}
