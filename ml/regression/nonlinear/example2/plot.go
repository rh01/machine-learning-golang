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

	df := dataframe.ReadCSV(f)

	yVals := df.Col("Sales").Float()
	radioVals := df.Col("Radio").Float()

	pts := make(plotter.XYZs, df.Nrow())
	ptsPred := make(plotter.XYZs, df.Nrow())

	// file
	for i, floatVal := range df.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = radioVals[i]
		pts[i].Z = yVals[i]
		ptsPred[i].X = floatVal
		pts[i].Y = radioVals[i]
		pts[i].Z = predict(floatVal, radioVals[i])

	}

	

	// new a plot
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	// p.Z.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Radius = vg.Points(3)

	

	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	// add
	p.Add(s, l)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "regression_line.pdf"); err != nil {
		log.Fatal(err)
	}

	//

}

func predict(tv, radio float64) float64 {
	return 2.93 + tv*0.05 + radio*0.18
}
