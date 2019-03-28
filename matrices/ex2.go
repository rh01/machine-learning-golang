package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {
	// using gonum.org/v1/gonum/floats
	// using gonum.org/v1/gonum/mat

	// vector := mat.NewVecDense(2, []float64{11.0, 2.5})

	// operator
	vA := []float64{11.0, 2.5, 2.0}
	vB := []float64{11.0, 2.5, 2.0}

	// dot
	v := floats.Dot(vA, vB)
	fmt.Println(v)

	// scale vector by scalar
	floats.Scale(2, vA)
	fmt.Println(vA)

	// 2-Norm
	nB := floats.Norm(vB, 2)
	fmt.Println(nB)


	// fmt.Printf("vector: %v", vector.T())
	// vector.AddScaledVec()

}
