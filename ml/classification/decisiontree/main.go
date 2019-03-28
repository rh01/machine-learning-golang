package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {
	ins, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(44111342)

	tree := trees.NewID3DecisionTree(0.6)

	// create cv
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(ins, tree, 5)
	if err != nil {
		log.Fatal(err)
	}

	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	fmt.Printf("\nAcuracy\n %0.2f (+/- %0.2f)]n", mean, stdev*2)
}
