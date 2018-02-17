package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func TrainModel() {
	// we open the csv file from the disk
	f, err := os.Open("./data/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// we create a new csv reader specifying
	// the number of columns it has
	salesData := csv.NewReader(f)
	salesData.FieldsPerRecord = 8

	// we read all the records
	records, err := salesData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//btc weighted price vs. x
	var r regression.Regression
	r.SetObserved("Price")
	r.SetVar(0, "Days")

	// Loop of records in the CSV, adding the training data to the regressionvalue.
	for i, record := range records {

		fmt.Println("record: ", record)
		// Skip the header.
		if i == 0 {
			continue
		}

		price, err := strconv.ParseFloat(records[i][7], 64)
		if err != nil {
			log.Fatal(err)
		}

		//use loop counter as independent variable (linear)
		grade, err := strconv.ParseFloat(strconv.Itoa(i), 64)

		if err != nil {
			log.Fatal(err)
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(price, []float64{grade}))
	}

	// Train/fit the regression model.
	r.Run()
	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}
