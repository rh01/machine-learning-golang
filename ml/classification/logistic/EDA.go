package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot/plotter"

	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("./clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	dfDescribe := df.Describe()
	fmt.Println(dfDescribe)

	for _, colName := range df.Names() {
		// create plotter.values 存储不同的数据源，用于画图
		plotterValues := make(plotter.Values, df.Nrow())
		// fill the data to plotterValue
		for idx, floatValue := range df.Col(colName).Float() {
			plotterValues[idx] = floatValue
		}

		// make a plt
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of %s", colName)

		// 创建直方图
		h, err := plotter.NewHist(plotterValues, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalize
		h.Normalize(1)

		// add h to plot
		p.Add(h)

		// save
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}

	}

}
