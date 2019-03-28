package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	trainingNum := (df.Nrow() * 4) / 5
	testNum := df.Nrow() / 5

	if trainingNum+testNum < df.Nrow() {
		trainingNum++
	}

	// 创建子集
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// 枚举所有的训练
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// 创建子df
	trainingDF := df.Subset(trainingIdx)
	testDF := df.Subset(testIdx)

	// 创建map
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// 创建对应的文件
	for idx, setName := range []string{"training.csv", "test.csv"} {
		// save the filted dataset file
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		w := bufio.NewWriter(f)

		// write the dataframe to a csv
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}

}
