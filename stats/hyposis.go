package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	observed := []float64{48, 32}
	expected := []float64{50, 50}

	// calculate the chisquare test statistic
	chiSquare := stat.ChiSquare(observed, expected)
	fmt.Println(chiSquare)

	
}
