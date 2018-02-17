package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
)

func CsvSplit(filename string) {

	f, err := os.Open("./data/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// we create a new csv reader specifying
	// the number of columns it has
	btcPriceData := csv.NewReader(f)
	btcPriceData.FieldsPerRecord = 8

	// we read all the records
	records, err := btcPriceData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// save the header
	header := records[0]

	// we have to shuffle the dataset before splitting
	// to avoid having ordered data
	// if the data is ordered, the data in the train set
	// and the one in the test set, can have different
	// behavior
	shuffled := make([][]string, len(records)-1)
	perm := rand.Perm(len(records) - 1)
	for i, v := range perm {
		shuffled[v] = records[i+1]
	}

	// split the training set
	trainingIdx := (len(shuffled)) * 4 / 5
	trainingSet := shuffled[1 : trainingIdx+1]

	// split the testing set
	testingSet := shuffled[trainingIdx+1:]

	// we write the splitted sets in separate files
	sets := map[string][][]string{
		"./data/training.csv": trainingSet,
		"./data/testing.csv":  testingSet,
	}

	for fn, dataset := range sets {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		out := csv.NewWriter(f)
		if err := out.Write(header); err != nil {
			log.Fatal(err)
		}

		if err := out.WriteAll(dataset); err != nil {
			log.Fatal(err)
		}
		out.Flush()
	}
}
