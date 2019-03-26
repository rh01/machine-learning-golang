package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type CSVRecords struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Variety     string
	ParseError  error
}

func main() {
	f, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal("File read failed.")
	}
	defer f.Close()

	r := csv.NewReader(f)
	// 设置每次只读一行
	r.FieldsPerRecord = -1
	// 设置一行的属性有五个
	// 用于处理属性不匹配
	r.FieldsPerRecord = 5

	var rawCSVData CSVRecords
	var CSVData []CSVRecords

	for {
		// read in a row, check if we are at the end of the file.
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		for idx, value := range record {
			if idx == 4 {
				// 解析 string
				if value == "" {
					log.Printf("Unexpect type in the column %d\n", idx)
					rawCSVData.ParseError = fmt.Errorf("Empty string value")
					break
				}

				// Add the value to the CSVRecord
				rawCSVData.Variety = value
				continue
			}

			// 否则，解析值为float64
			var floatValue float64

			// 如果不能解析为float， break
			if floatValue, err := strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Unexpect type in column %d\n", idx)
				rawCSVData.ParseError = fmt.Errorf("Could not parse float")
			}

			switch idx {
			case 0:
				rawCSVData.SepalLength = floatValue
			case 1:
				rawCSVData.SepalWidth = floatValue
			case 2:
				rawCSVData.PetalLength = floatValue
			case 3:
				rawCSVData.PetalWidth = floatValue
			}

		}

		// append the record to our dataset
		CSVData = append(CSVData, rawCSVData)
	}

}
