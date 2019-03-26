package main

import (
	"fmt"
	"log"

	gg "gorgonia.org/gorgonia"
)

func main() {
	g := gg.NewGraph()

	var x, y, z *gg.Node
	var err error

	// define the expression
	x = gg.NewScalar(g, gg.Float64, gg.WithName("x"))
	y = gg.NewScalar(g, gg.Float64, gg.WithName("y"))
	z, err = gg.Add(x, y)
	if err != nil {
		log.Fatal(err)
	}

	// create vm to compute.
	machine := gg.NewLispMachine(g)

	// initialize value the run.
	gg.Let(x, 2)
	gg.Let(y, 3)

	if machine.RunAll() != nil {
		log.Fatal(err)
	}

	log.Printf("data %v\n", z.Value())

	xgrad, err := x.Grad()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dz/dx: %v\n", xgrad)

	ygrad, err := y.Grad()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dz/dy: %v\n", ygrad)

}
