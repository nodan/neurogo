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

	g := gogame.NewGame()
	for !g.Finished() {
		c := g.Turn()
		s := n.Calculate(g.Board().Neural(c))
		for {
			xy := gogame.BestMove(s)
			if xy<0 {
				g.Pass()
				break
			}

			if g.Move(xy%g.Size(), xy/g.Size()) {
				break
			}
		}
	}

	fmt.Println(g.ShowAllPositions())

	for _, p := g.Positions() {
		c := g.Turn()
		b := g.Board().Neural(c)
		s := n.Calculate(b)
		gogame.Demote(s, p.Played())
		learn.Learn(n, b, s, 0.1)
	}
}
