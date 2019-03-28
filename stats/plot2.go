package main

import (
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

	// 画图
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	// 创建 箱
	w := vg.Points(50)

	irisDF := dataframe.ReadCSV(f)
	// 创建一个直方图
	for idx, colName := range irisDF.Names() {
		// 除去variety
		if colName != "variety" {
			// 创建 plotter.Values 并且使用df的对应列的值填充
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// 创建直方图
			h, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}

			// Normalize
			// h.Normalize(1)

			// 将直方图添加到plot
			p.Add(h)

		}

	}
	p.NominalX("sepal.length", "sepal.width", "petal.length", "petal.width")
	// save to png
	if err := p.Save(6*vg.Inch, 8*vg.Inch, "booxplots.png"); err != nil {
		log.Fatal(err)
	}

}
