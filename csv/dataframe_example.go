package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal("File read failed.")
	}
	defer f.Close()

	// r := csv.NewReader(f)

	// create dataframe
	df := dataframe.ReadCSV(f)
	fmt.Println(df)
	fmt.Println(df.Describe())

	// create the filter for dataframe
	filter := dataframe.F{
		Colname:    "variety",
		Comparator: "==",
		Comparando: "Virginica",
	}

	VirginicaDF := df.Filter(filter)
	if VirginicaDF.Err != nil {
		log.Fatal(df.Err)
	}

	// select operation
	VirginicaDF = df.Filter(filter).Select([]string{"sepal.length", "sepal.width"})

	fmt.Println(VirginicaDF)

	// select and limit first three rows
	VirginicaDF = df.Filter(filter).Select([]string{"sepal.length", "sepal.width"}).Subset([]int{0, 1, 2})

	fmt.Println(VirginicaDF)

	// df.Records()
}
