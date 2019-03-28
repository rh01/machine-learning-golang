package main

import (
	"fmt"
	"log"

	"github.com/sjwhitworth/golearn/evaluation"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/filters"
	"github.com/sjwhitworth/golearn/naive"
)

func main() {
	ins, err := base.ParseCSVToInstances("train.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	nb := naive.NewBernoulliNBClassifier()

	nb.Fit(convertToBinary(ins))

	testData, err := base.ParseCSVToTemplatedInstances("test.csv", true, ins)
	if err != nil {
		log.Fatal(err)
	}

	predictions, err := nb.Predict(convertToBinary(testData))
	if err != nil {
		log.Fatal(err)
	}

	cm, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		log.Fatal(err)
	}

	acc := evaluation.GetAccuracy(cm)
	fmt.Printf("acc : %0.2f\n", acc)

}

func convertToBinary(src base.FixedDataGrid) base.FixedDataGrid {
	b := filters.NewBinaryConvertFilter()
	attrs := base.NonClassAttributes(src)
	for _, a := range attrs {
		b.AddAttribute(a)
	}

	b.Train()
	ret := base.NewLazilyFilteredInstances(src, b)
	return ret
}
