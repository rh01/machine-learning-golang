package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal("File read failed.")
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Printf("csv read faild: %v", err)
	}

	for _, record := range records {
		index, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("%v", err)
		}
		fmt.Println(index)
	}

}
