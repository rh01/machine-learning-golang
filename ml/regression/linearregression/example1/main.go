package main

import (
	"fmt"
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

	// 简单的描述统计
	advertSummary := df.Describe()

	// 打印总结性的统计
	fmt.Println(advertSummary)

	

}
