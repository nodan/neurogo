package main

import (
	"fmt"
	"math"
	"math/rand"
	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"github.com/NOX73/go-neural/persist"
)

func main() {
	// 3 layers: inputs, processing, outputs
	n := neural.NewNetwork(3, []int{9,81,9})

	// Randomize sypaseses weights
	n.RandomizeSynapses()

	// Learning speed [0..1]
	for i:=0; i<1000000; i++ {
		r := []float64{ rand.Float64(), rand.Float64(), rand.Float64() }
		s := []float64{ r[1]/2, r[2]/3, r[0]/4 }
		learn.Learn(n, r, s, 0.1)
	}

	for i:=0; i<10; i++ {
		r := []float64{ rand.Float64(), rand.Float64(), rand.Float64() }
		s := n.Calculate(r)
		s[0] -= r[1]/2
		s[1] -= r[2]/3
		s[2] -= r[0]/4
		fmt.Println(r, s, math.Sqrt(s[0]*s[0]+s[1]*s[1]+s[2]*s[2]/3))
	}

	persist.ToFile("network.json", n);
}
