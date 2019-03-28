package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sajari/regression"
)

type ModelInfo struct {
	Intercept    float64           `json:"intercept"`
	Coefficients []CoefficientInfo `json:"coefficients"`
}

type CoefficientInfo struct {
	Name        string  `json:"name"`
	Coefficient float64 `json:"coefficient"`
}

func main() {

	initDir := flag.String("initDir", "", "This Directory containinng the training data")
	outDir := flag.String("outDir", "", "The output dir")

	flag.Parse()

	f, err := os.Open("./training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create a new csv reeader
	reader := csv.NewReader(f)

	// read in all of the csv records
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 尝试创建一个模型 Sales(Y) = TV(X) * a + b
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")

	// 循环
	for i, record := range trainingData {
		// skip the head
		if i == 0 {
			continue
		}

		// 解析TV
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// add these point to the regression value
		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))
	}

	r.Run()

	fmt.Printf("\n regreession formula: \n%v\n\n", r.Formula)

	// 评估模型
	// 参考第三章
	f, err = os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader = csv.NewReader(f)
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var mAE float64
	for i, record := range testData {

		if i == 0 {
			continue
		}

		// read in the observed and preedicted value
		observedValue, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
			continue
		}

		// read in the observed and preedicted value
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
			continue
		}

		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
			continue
		}

		yPredict, err := r.Predict([]float64{tvVal, radioVal})
		if err != nil {
			log.Fatal(err)
		}

		mAE += math.Abs(observedValue-yPredict) / float64(len(testData))

	}

	fmt.Printf("\nMAE = %0.2f\n", mAE)
	// fmt.Printf("\nMSE = %0.2f\n", mSE)

	// save model!!!

	modelInfo := ModelInfo{
		Intercept: r.Coeff(0),
		Coefficients: []CoefficientInfo{
			CoefficientInfo{
				Name:        "bmi",
				Coefficient: r.Coeff(1),
			},
			CoefficientInfo{
				Name:        "ltg",
				Coefficient: r.Coeff(2),
			},
		},
	}

	b, err := json.MarshalIndent(modelInfo, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filepath.Join(".", "model.json"), b, 0644); err != nil {
		log.Fatal(err)
	}

}
