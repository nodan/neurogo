package main

import (
	"fmt"
	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/learn"
	"net/http"
	"neurogo/gogame"
)

// 3 layers: inputs, processing, outputs
var n = neural.NewNetwork(9, []int{9, 81, 9})

// handle http requests
func boardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<html>
<head>
<meta name="Content-Type" content="text/html; charset=UTF-8" />`)

	// URL: "/<color to move><board>"
	// with X=black, O=white, .=empty
	c, g := gogame.Parse(r.URL.Path[1:])
	b := g.Neural(c)
	s := n.Calculate(b)

	fmt.Fprintf(w, `
</head>
<body>
<p><table>`)
	// print the go board as table
	for y := 0; y < gogame.Size; y++ {
		fmt.Fprintf(w, "<tr height=\"20px\">")
		for x := 0; x < gogame.Size; x++ {
			xy := gogame.Xy(x, y)
			// show shades of red...green for 0...1
			fmt.Fprintf(w, "<td align=\"center\" width=\"20px\" style=\"background-color:#%02x%02xbf\">",
				int(0xbf+0x40*(1.0-s[gogame.Xy(x, y)])),
				int(0xbf+0x40*s[xy]))
			switch g[xy] {
			case gogame.Black:
				fmt.Fprintf(w, "X")
			case gogame.White:
				fmt.Fprintf(w, "O")
			}
			fmt.Fprintf(w, "</td>")
		}
		fmt.Fprintf(w, "</tr>")
	}
	fmt.Fprintf(w, "</table>")

	fmt.Fprintf(w, `
</body>
</html>`)
}

func main() {
	// install the http handler
	http.HandleFunc("/", boardHandler)

	// start with a random network
	n.RandomizeSynapses()

	// learn that 1,1 is a good move
	grid := &gogame.Grid{}
	b := grid.Neural(gogame.Black)
	s := append(b, 1)[0:9]
	s[4] = 1
	for i := 0; i < 1000; i++ {
		learn.Learn(n, b, s, 1)
	}

	// for i:=0; i < 10; i++ {
	// 	g := playAiSoloGame(n)
	// 	learnFrom(g, n)

	// 	// fmt.Println(g.ShowGame())
	// 	// fmt.Println("score", g.Board().Score())
	// }

	// g := playAiSoloGame(n)
	// fmt.Println(g.ShowGame())
	// fmt.Printf("Score %v after %v moves\n", g.Board().Score(), len(g.Positions()))

	// start the http server
	http.ListenAndServe(":8080", nil)
}

// let the neural net play a game against itself
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

// find the best move, i.e. the highest value
func bestMove(s []float64) int {
	n := gogame.Size
	rc := -1
	for xy := 0; xy < n*n; xy++ {
		if s[xy] >= 0 && (rc < 0 || s[xy] > s[rc]) {
			rc = xy
		}
	}

	if rc >= 0 {
		s[rc] = -1
	}
	return rc
}

// learn from a given game by promoting/demoting moves
// of the winning/losing color
func learnFrom(g *gogame.Game, n *neural.Network) {
	pn := g.Positions()
	score := g.Score()
	if score == 0 {
		return
	}

	var lost gogame.Color
	lost = gogame.Black
	if score > 0 {
		lost = gogame.White
	}

	l := len(pn)
	// if l > 6 {
	//	l -= 6
	// } else {
	//	l = 0
	// }
	l = 0

	for _, p := range pn[l:] {
		c := g.Turn()
		b := g.Board().Neural(c)
		s := n.Calculate(b)
		if c == lost && p.Move.MoveType != gogame.Pass {
			demote(s, gogame.Xy(p.Move.X, p.Move.Y))
			learn.Learn(n, b, s, 0.2)
		} else if c != lost && p.Move.MoveType != gogame.Pass {
			s[gogame.Xy(p.Move.X, p.Move.Y)] = 1
			learn.Learn(n, b, s, 0.1)
		}
	}
}

// demote a given move either below the next best move or
// simply to 0
func demote(s []float64, xy int) {
	//n := gogame.Size
	// find the next best move
	nb := 1
	// for nxy := 0; nxy < n*n; nxy++ {
	//	if s[nxy]>=0 && (nb<0 || s[nxy]>s[nb]) && s[nxy]<s[xy] {
	//		nb = nxy
	//	}
	// }

	if nb >= 0 {
		// demote xy
		s[xy] = 0 // s[nb]*0.9
	}
}
