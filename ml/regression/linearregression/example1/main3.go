package main

import (
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot/plotter"

	"github.com/kniren/gota/dataframe"
)

func main() {
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 从csv文件中创建一个df
	df := dataframe.ReadCSV(f)

	// 提取target列
	yValsss := df.Col("Sales").Float()

	for _, colName := range df.Names() {
		// make plotter.values
		pts := make(plotter.XYs, df.Nrow())

		// file pts with data
		// 对每一列进行遍历，构造每个属性与目标的pair
		for i, floatVal := range df.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yValsss[i]
		}
		// 创建plot
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		// p.Title.Text()
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		// 圆点的直径
		s.GlyphStyle.Radius = vg.Points(3)

		p.Add(s)

		// save
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scat.png"); err != nil {
			log.Fatal(err)
		}

	}

}
