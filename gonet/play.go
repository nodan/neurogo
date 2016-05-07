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
	lastMovePass := false

	g := gogame.NewGame()
	game: for !g.Finished() {
		c := g.Turn()
		s := n.Calculate(g.Board().Neural(c))
		for {
			xy := bestMove(s)
			if xy < 0 {
				g.Pass()
				if lastMovePass {
					break game
				}
				lastMovePass = true
				break
			}

			if g.Move(xy%g.Size(), xy/g.Size()) {
				lastMovePass = false
				break
			}
		}
	}

	fmt.Println(g.ShowGame())

	// for _, p := g.Positions() {
	//	c := g.Turn()
	//	b := g.Board().Neural(c)
	//	s := n.Calculate(b)
	//	demote(s, p.Played())
	//	learn.Learn(n, b, s, 0.1)
	// }
}


func bestMove(s []float64) int {
	n := gogame.Size
	rc := -1
	for xy := 0; xy < n*n; xy++ {
		if s[xy]>=0 && (rc<0 || s[xy]>s[rc]) {
			rc = xy
		}
	}

	if rc>=0 {
		s[rc] = -1
	}
	return rc
}

func demote(s []float64, xy int) {
	n := gogame.Size
	// find the next best move
	nb := -1
	for nxy := 0; nxy < n*n; nxy++ {
		if s[nxy]>=0 && (nb<0 || s[nxy]>s[nb]) && s[nxy]<s[xy] {
			nb = nxy
		}
	}

	if nb>=0 {
		// demote xy
		s[xy] = s[nb]*0.9
	}
}
