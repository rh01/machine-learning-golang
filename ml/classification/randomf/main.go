package main

import (
	"fmt"
	"log"
	"math"

	"github.com/sjwhitworth/golearn/evaluation"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
)

func main() {
	ins, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	rf := ensemble.NewRandomForest(10, 4)

	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(ins, rf, 5)
	if err != nil {
		log.Fatal(err)
	}

	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	fmt.Printf("\nAcuracy\n %0.2f (+/- %0.2f)]n", mean, stdev*2)

}
