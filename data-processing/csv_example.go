package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var intMax = math.MinInt64

func main() {
	// ioutil.ReadFile("")
	// Open the CSV
	f, err := os.Open("myfile.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read in the csv records
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		// Parse the integer Value
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		// Replace thee maximum valuee if appropriate
		if intVal > intMax {
			intMax = intVal
		}

		fmt.Println(intMax)
	}

}
