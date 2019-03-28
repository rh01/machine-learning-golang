package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	// df.Records()

	var observed []int
	var predicted []int
	line := 1

	for {

		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// skip header
		if line == 1 {
			line++
			continue
		}

		// read in the observed and preedicted value
		observedValue, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
			continue
		}

		predictValue, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
			continue
		}

		observed = append(observed, observedValue)
		predicted = append(predicted, predictValue)
		line++
	}

	var turePosNeg int

	for iidx, oVal := range observed {
		if oVal == predicted[iidx] {
			turePosNeg++
		}
	}

	// 计算准确度
	accuracy := float64(turePosNeg) / float64(len(observed))
	fmt.Printf("\nAcc = %0.2f\n\n", accuracy)

	// 下面计算召回率和准确率
	classes := []int{0, 1, 2}

	// loop each class
	for _, class := range classes {
		var truePos int
		var falsePos int
		var falseNeg int

		for idx, oVal := range observed {
			switch oVal {
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}
				falseNeg++
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}
		precision := float64(truePos) / float64(truePos+falsePos)
		recall := float64(truePos) / float64(truePos+falseNeg)
		fmt.Printf("\n precision (class %d)=%0.2f", class, precision)
		fmt.Printf("\n recall (class %d)=%0.2f", class, recall)
	}

}
