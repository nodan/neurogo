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
		for {
			xy := g.BestMove(s)
			if xy<0 {
				fmt.Println("pass", c)
				break
			}

			fmt.Println("best move", xy, c)
			if g.MakeMove(xy, c)!=nil {
				break
			}
		}
		fmt.Println(g)
		c = gogame.Invert(c)
	}
}
