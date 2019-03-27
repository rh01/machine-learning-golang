package main

import (
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot/plotter"

	"gonum.org/v1/plot"
)

func main() {
	// from sratch
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Logistic Function"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x)"

	// 创建画图函数
	logisticFunction := plotter.NewFunction(logistic)
	logisticFunction.Color = color.RGBA{B: 255, A: 255}

	p.Add(logisticFunction)

	p.X.Min = -10
	p.X.Max = 10
	p.Y.Min = -0.1
	p.Y.Max = 1.1

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "logistic_regression.png"); err != nil {
		log.Fatal(err)
	}

}

func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-1*x))
}
