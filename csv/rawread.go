package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

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
	

	var rawCSVData [][]string

	for {
		// read in a row, check if we are at the end of the file.
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
			continue
		}

		// append the record to our dataset
		rawCSVData = append(rawCSVData, record)
	}

}
