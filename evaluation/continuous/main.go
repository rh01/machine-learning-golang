package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

func main() {
	f, err := os.Open("./continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	// df.Records()

	var observed []float64
	var predicted []float64
	line := 1

	for {

		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// skip header
		if line == 1 {
			line++
			continue
		}

		// read in the observed and preedicted value
		observedValue, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
			continue
		}

		predictValue, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
			continue
		}

		observed = append(observed, observedValue)
		predicted = append(predicted, predictValue)
		line++
	}

	var mAE float64
	var mSE float64

	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal-predicted[idx], 2) / float64(len(observed))
	}

	fmt.Printf("\nMAE = %0.2f\n", mAE)
	fmt.Printf("\nMSE = %0.2f\n", mSE)

	// utilize the stat lib
	rSqure := stat.RSquaredFrom(observed, predicted, nil)
	fmt.Println(rSqure)

}
