package main

import (
	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"fmt"
)

func main() {
	n := neural.NewNetwork(2, []int{2,1})
	n.RandomizeSynapses()
	result := n.Calculate([]float64{0.6, 0.7})
	fmt.Printf("Result: %v\n", result)

	learn.Learn(n, []float64{0, 0}, []float64{0}, 1)
	learn.Learn(n, []float64{0.25, 0.25}, []float64{0.5}, 1)
	learn.Learn(n, []float64{0.5, 0.5}, []float64{1}, 1)
	learn.Learn(n, []float64{0.75, 0.75}, []float64{1.5}, 1)
	learn.Learn(n, []float64{1, 1}, []float64{2}, 1)

	result = n.Calculate([]float64{0.6, 0.7})
	fmt.Printf("Result after learning: %v\n", result)

	e := learn.Evaluation(n, []float64{0.6, 0.7}, []float64{10})
	fmt.Printf("Eval: %v\n", e)
}
