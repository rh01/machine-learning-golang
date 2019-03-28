package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {
	distance := floats.Distance([]float64{1, 2}, []float64{2, 3}, 2)
	fmt.Println(distance)
}
