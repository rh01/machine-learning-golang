package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("./clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create dataframe of f
	loanDF := dataframe.ReadCSV(f)

	// create training idx and test dataset idx
	trainingNum := (4 * loanDF.Nrow()) / 5
	testNum := loanDF.Nrow() / 5

	if trainingNum+testNum < loanDF.Nrow() {
		trainingNum++
	}

	// create subset
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// 初始化两个序列
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	for i := 0; i < testNum; i++ {
		testIdx[i] = i + trainingNum
	}

	trainingDF := loanDF.Subset(trainingIdx)
	testDF := loanDF.Subset(testIdx)

	// 创建一个map，用来吧数据写入文件中
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// 创建对应的文件
	for index, setName := range []string{"train.csv", "test.csv"} {
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		// 创建bufer writer
		w := bufio.NewWriter(f)

		// 将dataframe写出到csv文件格式
		if err = setMap[index].WriteCSV(w); err != nil {
			log.Fatal(err)
		}

	}

}
