package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

type centroid []float64

func main() {
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	irisDF := dataframe.ReadCSV(f)

	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	centroids := make(map[string]centroid)

	for _, species := range speciesNames {
		// filter the origin dataset
		fileter := dataframe.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(fileter)

		// calculate the mean of feature
		summeryDF := filtered.Describe()

		var c centroid
		for _, feature := range summeryDF.Names() {
			// skip
			if feature == "column" || feature == "species" {
				continue
			}
			c = append(c, summeryDF.Col(feature).Float()[0])
		}

		// add
		centroids[species] = c

	}

	for _, spec := range speciesNames {
		fmt.Printf("%s centroid: %v\n", spec, centroids[spec])
	}

}
