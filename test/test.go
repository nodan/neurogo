package main

import (
	"math/rand"
	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"fmt"
)

func main() {
	n := neural.NewNetwork(2, []int{2, 1})
	n.RandomizeSynapses()
	result := n.Calculate([]float64{0.6, 0.7})
	fmt.Printf("Result: %v\n", result)

	for i := 0; i < 1000000; i++ {
		x := rand.Float64()
		y := rand.Float64()
		learn.Learn(n, []float64{x, y}, []float64{x * y}, 1)
	}

	for i := 0; i < 10; i++ {
		x := rand.Float64()
		y := rand.Float64()
		z := n.Calculate([]float64{x, y})
		fmt.Printf("%v * %v = %v <%v> <%v>\n", x, y, z,
			learn.Evaluation(n, []float64{x, y}, []float64{x * y}),
			z[0]-x*y)
	}

}
