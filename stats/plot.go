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
	f, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal("File read failed.")
	}
	defer f.Close()

	irisDF := dataframe.ReadCSV(f)
	// 创建一个直方图
	for _, colName := range irisDF.Names() {
		// 除去variety
		if colName != "variety" {
			// 创建 plotter.Values 并且使用df的对应列的值填充
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// 画图
			p, err := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			// 创建直方图
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			// Normalize
			h.Normalize(1)

			// 将直方图添加到plot
			p.Add(h)

			// save to png
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}

		}

	}

}
