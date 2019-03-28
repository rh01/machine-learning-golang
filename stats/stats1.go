package main

import (
	"fmt"
	"log"
	"os"

	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"

	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal("File read failed.")
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	ser := df.Col("sepal.length").Float()
	// 中位数
	median, err := stats.Median(ser)
	if err != nil {
		log.Fatal(err)
	}

	// 众数和众数的数目
	mode, count := stat.Mode(ser, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 平均数
	mean, err := stats.Mean(ser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Midian: %0.2f\n", median)
	fmt.Printf("Mode: %0.2f\n", mode)
	fmt.Printf("Mode Count: %0.2f\n", count)
	fmt.Printf("Mean: %0.2f\n", mean)
}
