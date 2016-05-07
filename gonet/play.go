package main

import (
	"fmt"
	"neurogo/gogame"
	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
//	"github.com/NOX73/go-neural/persist"
)

func main() {
	// 3 layers: inputs, processing, outputs
	n := neural.NewNetwork(9, []int{9, 81, 9})
	n.RandomizeSynapses()

	for i:=0; i<1000; i++ {
		g := playAiSoloGame(n)
		learnFrom(g, n)

		fmt.Println(g.ShowGame())
		fmt.Println("score", g.Board().Score())
	}
}

func playAiSoloGame(n *neural.Network) *gogame.Game {
	g := gogame.NewGame()
	for !g.Finished() {
		c := g.Turn()
		s := n.Calculate(g.Board().Neural(c))
		for {
			xy := bestMove(s)
			if xy < 0 {
				g.Pass()
				break
			}
			if g.Move(xy%g.Size(), xy/g.Size()) {
				break
			}
		}
	}

	return g
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

func learnFrom(g *gogame.Game, n *neural.Network) {
	pn := g.Positions()
	score := g.Score()-8
	if score == 0 {
		return
	}

	var lost gogame.Color
	lost = gogame.Black
	if score>0 {
		lost = gogame.White
	}

	l := len(pn)
	if l>20 {
		l -= 20
	} else {
		l = 0
	}

	for _, p := range pn[l:] {
		c := g.Turn()
		b := g.Board().Neural(c)
		s := n.Calculate(b)
		if c==lost && p.Move.MoveType!=gogame.Pass {
			demote(s, gogame.Xy(p.Move.X, p.Move.Y))
			learn.Learn(n, b, s, 0.1)
		}
	}
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
