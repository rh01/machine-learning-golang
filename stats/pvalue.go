package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	observed := []float64{260.0, 135.0, 105.0}

	totalObserved := 500.0

	expected := []float64{totalObserved * 0.60, totalObserved * 0.25, totalObserved * 0.15}

	// calculate the chisquare test statistic
	chiSquare := stat.ChiSquare(observed, expected)
	fmt.Println(chiSquare)

	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}

	// 计算 p-value
	pValue := chiDist.Prob(chiDist)

	fmt.Println(pValue)

}
