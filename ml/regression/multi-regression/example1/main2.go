package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot"

	"gonum.org/v1/plot/plotter"

	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	// 为每一列提供一个直方图
	for _, colName := range df.Names() {
		// 创建plotter.Value
		plotValues := make(plotter.Values, df.Nrow())
		for i, floatVal := range df.Col(colName).Float() {
			plotValues[i] = floatVal
		}

		// make a plot
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of %s", colName)

		// 创建呢一个直方图
		h, err := plotter.NewHist(plotValues, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalize -> 使得总面积为1
		h.Normalize(1)

		// 增加plottt
		p.Add(h)

		// save the result
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_histt.png"); err != nil {
			log.Fatal(err)
		}

	}

}
