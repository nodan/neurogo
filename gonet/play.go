package main

import (
	"fmt"
	"neurogo/gogame"
	"github.com/NOX73/go-neural"
//	"github.com/NOX73/go-neural/learn"
//	"github.com/NOX73/go-neural/persist"
)

func main() {
	// 3 layers: inputs, processing, outputs
	n := neural.NewNetwork(9, []int{9, 81, 9})
	n.RandomizeSynapses()

	var g gogame.Grid
	var c gogame.Color
	c = gogame.Black
	for !g.Finished() {
		s := n.Calculate(g.Neural(c))
		fmt.Println(s)
		if g.MakeMove(g.BestMove(s), c)==nil {
			fmt.Println("illegal move")
			break
		}
		fmt.Println(g)
		c = gogame.Invert(c)
	}
}
