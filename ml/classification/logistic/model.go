package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gonum/matrix/mat64"
)

func main() {
	f, err := os.Open("./train.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create a new csv reeader
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	// read all records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 因为有常数/截距
	featureData := make([]float64, 2*len(records))
	labels := make([]float64, len(records))

	// feature matrix， 我们要将featureData换成featureMatrix
	// featureMatrix = make(mat64.Dense)
	var featureIndex int

	for idx, record := range records {
		// skip header
		if idx == 0 {
			continue
		}

		// add
		// 解析TV
		scoreVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		featureData[featureIndex] = scoreVal

		featureData[featureIndex+1] = 1.0

		featureIndex += 2

		// add label

		label, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		labels[idx-1] = label

	}

	// 将feature转成matrix
	features := mat64.NewDense(len(records), 2, featureData)

	//train
	weights := logisticRegress(features, labels, 1000, 0.3)

	formula := "p = 1 / (1 + exp( - m1 * FICO.score - m2) "
	fmt.Printf("\n%s\n\nm1 = %0.2f\nm2 = %0.2f\n\n", formula, weights[0], weights[1])
}

// 对于给定的训练数据，建立逻辑回归模型
//
// features: mat64.Dense Matrix 类型的指针, 该matrix包含了一行
// labels: []float64类型，对应featture的label
// numStep: 最大训练步数
// learningRate: 表示学习率
// 返回值:
func logisticRegress(features *mat64.Dense, labels []float64, numStep int, learningRate float64) []float64 {

	// 初始化随机的权重
	_, numWeights := features.Dims()
	weights := make([]float64, numWeights)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for idx, _ := range weights {
		weights[idx] = r.Float64()
	}

	// 迭代
	for i := 0; i < numStep; i++ {

		// if i == 0 {}
		// 初始化 错误
		var sumError float64

		// make prediction for each lable and accumatelu error
		for idx, label := range labels {
			// 通过 idx 得到对应的featue值
			featureRow := mat64.Row(nil, idx, features)

			// 计算logistttic值
			pred := logistic(featureRow[0]*weights[0] + featureRow[1]*weights[1])
			predError := label - pred
			sumError += math.Pow(predError, 2)

			// update the feature weight
			for j := 0; j < len(featureRow); j++ {
				weights[j] += learningRate * predError * pred * (1 - pred) * featureRow[j]
			}
		}
	}

	return weights

}

func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-1*x))
}
