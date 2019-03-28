package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	vA := mat.NewVecDense(3, []float64{11.0, 2.0, 3.5})
	vB := mat.NewVecDense(3, []float64{11.0, 2.0, 3.5})

	v := mat.Dot(vA, vB)
	fmt.Println(v)

	// ??
	vA.ScaleVec(1.5, vA)

	v = mat.Norm(vA, 2)
	fmt.Println(v)


}
