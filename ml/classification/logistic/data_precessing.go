package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	scoreMax = 830.0
	scoreMin = 640.0
)

func main() {
	f, err := os.Open("./loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 创建一个csv reader，并且指定每行读取的属性数目
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 创建输出文件
	f, err = os.Create("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	// 声明和初始化两个数据

	for idx, record := range records {
		if idx == 0 {
			if err := w.Write(record); err != nil {
				log.Fatal(err)
			}
			continue
		}

		// 初始化一个数组，用于保存我们解析的数据
		outRecord := make([]string, 2)

		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 正规化
		outRecord[0] = strconv.FormatFloat((score-scoreMin)/(scoreMax-scoreMin), 'f', 4, 64)
		if err != nil {
			log.Fatal(err)
		}

		// rate, err := strconv.ParseFloat(strings.Split(record[1], "%")[0], 64)
		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}
		if rate <= 12.0 {
			outRecord[1] = "1.0"
			// write the record to output fiile
			if err := w.Write(outRecord); err != nil {
				log.Fatal(err)
			}
			continue
		}
		outRecord[1] = "0.0"
		if err := w.Write(outRecord); err != nil {
			log.Fatal(err)
		}
		// continue

	}

	// wrrite any buffer data to uderlay writer
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}
