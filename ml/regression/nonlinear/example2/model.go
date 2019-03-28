package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/berkmancenter/ridge"
	"github.com/gonum/matrix/mat64"
)

func main() {
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

	// 构建，然后将featureedData转成gonum.matrics
	featuredData := make([]float64, 4*len(trainingData))
	yData := make([]float64, len(trainingData))

	var featuredIndex int
	var yIndex int

	// 循环
	for i, record := range trainingData {
		// skip the head
		if i == 0 {
			continue
		}

		for i, val := range record {
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			if i < 3 {
				// add an intercept to the model
				if i == 0 {
					featuredData[featuredIndex] = 1
					featuredIndex++
				}
				featuredData[featuredIndex] = valParsed
				featuredIndex++
			}
			if i == 3 {
				// add the float to the slice of y floats
				yData[yIndex] = valParsed
				yIndex++
			}
		}
	}

	features := mat64.NewDense(len(trainingData), 4, featuredData)
	y := mat64.NewVector(len(trainingData), yData)

	// create ridgeregressionn
	r := ridge.New(features, y, 1.0)
	r.Regress()

	c1 := r.Coefficients.At(0, 0)
	c2 := r.Coefficients.At(1, 0)
	c3 := r.Coefficients.At(2, 0)
	c4 := r.Coefficients.At(3, 0)

	fmt.Printf("\n regreession formula: \ny = %0.3f + %0.3f TV + %0.3f Radio + %0.3f Newspaper\n\n",
		c1, c2, c3, c4)

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

		newspaperVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
			continue
		}

		yPredict := predict(tvVal, radioVal, newspaperVal)

		mAE += math.Abs(observedValue-yPredict) / float64(len(testData))

	}

	fmt.Printf("\nMAE = %0.2f\n", mAE)
	// fmt.Printf("\nMSE = %0.2f\n", mSE)

	

}

func predict(tv float64, radio float64, newspaper float64) float64 {
	return 3.038 + 0.047*tv + 0.177*radio + 0.001*newspaper
}
