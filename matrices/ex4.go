package main

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
)

func main() {
	components := []float64{1.5, 2.3, -5.7, -2.4, 7.3, 3.0}
	// components2 := []float64{1, 2, -5, -2, 7, 3}
	matrix1 := mat.NewDense(2, 3, components)
	// matrix2 := mat.NewDense(2, 3, components2)
	fmt.Println(matrix1.At(1, 2))

	fmt.Println(matrix1.Dims())
	// matrix1.Sub(matrix1, matrix2)
	// fmt.Println(matrix1)

	matrix1.SetCol(1, []float64{2, 3})
	fmt.Println(matrix1)

	v2 := matrix1.ColView(2)
	fmt.Println(v2)

	// 创建两个相同大小的矩阵
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})

	c := mat.NewDense(3, 2, []float64{3, 2, 1, 4, 0, 8})

	// add a and b
	d := mat.NewDense(3, 3, nil)               // 声明3x3的矩阵
	d.Add(a, b)                                // a + b -> d
	fd := mat.Formatted(d, mat.Prefix("    ")) // 返回一个fmt.Formatter
	fmt.Printf("d = a + b = %0.4v\n", fd)

	// 矩阵乘法， axc
	f := mat.NewDense(3, 2, nil)
	f.Mul(a, c)
	fd = mat.Formatted(f, mat.Prefix("    ")) // 返回一个fmt.Formatter
	fmt.Printf("f = a x b = %0.4v\n", fd)

	// raising a matrix to a power
	g := mat.NewDense(3, 3, nil)
	g.Pow(a, 5)
	fd = mat.Formatted(g, mat.Prefix("    ")) // 返回一个fmt.Formatter
	fmt.Printf("g = a^5 = %0.4v\n", fd)

	// apply a function to each of the elements of a
	h := mat.NewDense(3, 3, nil)
	h.Apply(func(_, _ int, v float64) float64 {
		return math.Sqrt(v)
	}, a)
	fd = mat.Formatted(h, mat.Prefix("    ")) // 返回一个fmt.Formatter
	fmt.Printf("g = sqrt(a) = %0.4v\n", fd)

	// 矩阵转置
	// i := mat.NewDense(3, 3, nil)
	// transposed := mat.Transpose{a}
	// transposed.T()
	fd = mat.Formatted(a, mat.Prefix("    ")) // 返回一个fmt.Formatter
	fmt.Printf("g = sqrt(a) = %0.4v\n", fd)

	// a.T()
	fd = mat.Formatted(a.T(), mat.Prefix("    ")) // 返回一个fmt.Formatter
	fmt.Printf("g = sqrt(a) = %0.4v\n", fd)

	// 行列式
	v := mat.Det(a)
	fmt.Println(v)

	// 矩阵的逆
	i := mat.NewDense(3, 3, nil)
	if err := i.Inverse(a); err != nil {
		log.Fatal(err)
	}
	mat.Formatted(i, mat.Prefix("    "))
	fmt.Printf("innverse of a= %0.4v", i)

}
