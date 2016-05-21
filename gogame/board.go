package gogame

const (
	// board size
	n    = 3
	Size = n

	// Colors
	Empty = 0
	Black = 1
	White = 2
)

type Color byte

// the board
type Grid [n * n]Color

// transform a board coordinate into a linear array index
func Xy(x, y int) int {
	return n*y + x
}

// Invert a Color Black<->White
func Invert(c Color) Color {
	switch c {
	case Black:
		return White
	case White:
		return Black
	default:
		return Empty
	}
}

// find neighbors of a point
func neighbors(xy int) []int {
	rc := make([]int, 0, 4)
	if xy%n >= 1 {
		rc = append(rc, xy-1)
	}

	if xy%n+1 != n {
		rc = append(rc, xy+1)
	}

	if xy >= n {
		rc = append(rc, xy-n)
	}

	if xy+n < n*n {
		rc = append(rc, xy+n)
	}

	return rc
}

// find up to max liberties of a chain
func (g *Grid) findLiberties(xy int, max int) int {
	libs := 0
	c := g[xy]
	opposite := Invert(c)
	g[xy] = opposite // don't look here again

	// look at the neighbors
	for _, nxy := range neighbors(xy) {
		switch g[nxy] {
		case Empty:
			// count liberty
			libs += 1
			g[nxy] = opposite // don't look here again

			// count up to max libs
			if libs >= max {
				return libs
			}
		case c:
			// recursively count liberties of neighbor
			libs += g.findLiberties(nxy, max-libs)

			// count up to max libs
			if libs >= max {
				return libs
			}
		}
	}

	// the liberties found so far
	return libs
}

// find up to max liberties of a chain
func (g *Grid) liberties(xy int, max int) int {
	t := *g
	return t.findLiberties(xy, max)
}

// remove a chain of stones
func (g *Grid) remove(xy int) *Grid {
	c := g[xy]
	g[xy] = Empty
	for _, nxy := range neighbors(xy) {
		if g[nxy] == c {
			g.remove(nxy)
		}
	}

	return g
}

// count liberties of a chain of stones
func (counter *Grid) count(xy int, g *Grid) *Grid {
	if c := g[xy]; c != Empty {
		g[xy] = Empty
		counter[xy]++
		for _, nxy := range neighbors(xy) {
			if g[nxy] == c {
				counter.count(nxy, g)
			}
		}
	}
	return counter
}

// convert a string into a grid
func Parse(s string) (Color, *Grid) {
	c := Color(Black)
	g := Grid{}

	if len(s) > 0 && s[0:1] == "O" {
		c = White
	}

	for xy := 0; xy+1 < len(s); xy++ {
		switch s[xy+1 : xy+2] {
		case "X":
			g[xy] = Black
		case "O":
			g[xy] = White
		}
	}

	return c, &g
}

// convert a grid into a string
func (g *Grid) String(c Color) string {
	s := ""
	if c == Black {
		s += "X"
	} else {
		s += "O"
	}

	for xy := 0; xy < n*n; xy++ {
		switch g[xy] {
		case Black:
			s += "X"
		case White:
			s += "O"
		default:
			s += "."
		}
	}

	return s
}

// play a move
func (g *Grid) MakeMove(xy int, c Color) *Grid {
	if g[xy] != Empty {
		// don't play on non-Empty points
		return nil
	}

	// play a move
	g[xy] = c

	// check neighbors
	for _, nxy := range neighbors(xy) {
		t := *g
		if t[nxy] == Invert(c) && t.findLiberties(nxy, 1) == 0 {
			// remove captured stones
			g.remove(nxy)
		}
	}

	// check liberties of the move played
	t := *g
	if t.findLiberties(xy, 1) == 0 {
		// undo the move
		g[xy] = Empty

		// illegal move, no liberties
		return nil
	}

	return g
}

// check if the game is finished in the sense of there not being two adjacent Empty points and
// every group having exactly two liberties
func (g *Grid) Finished() bool {
	var c Grid
	for xy := 0; xy < n*n; xy++ {
		// find Empty points
		if g[xy] == Empty {
			t := *g
			nl := neighbors(xy)
			for _, nxy := range nl {
				// same Color for all adjacent points
				if g[nxy] == Empty || g[nxy] != g[nl[0]] {
					return false
				}
			}

			for _, nxy := range nl {
				// count liberties
				c.count(nxy, &t)
				if c[nxy] > 2 { // abort early
					return false
				}
			}
		}
	}

	// find non-Empty points
	for xy := 0; xy < n*n; xy++ {
		// all chains must have 2 liberties
		if g[xy] != Empty && c[xy] != 2 {
			return false
		}
	}

	return true
}

// determine the score by counting black and white stones
// as well as empty points surrounded by one color
func (g *Grid) Score() int {
	score := 0
	for xy := 0; xy < n*n; xy++ {
		switch g[xy] {
		case Black:
			score++
		case White:
			score--
		case Empty:
			bn := false
			wn := false
			for _, nxy := range neighbors(xy) {
				switch g[nxy] {
				case Black:
					bn = true
				case White:
					wn = true
				}
			}
			if bn && !wn {
				score++
			} else if !bn && wn {
				score--
			}
		}
	}
	return score
}

// determine if the position is legal, i.e. all stones having
// at least one liberty
func (g *Grid) Legal() bool {
	for xy := 0; xy < n*n; xy++ {
		if g[xy] != Empty && g.liberties(xy, 1) == 0 {
			return false
		}
	}
	return true
}

// convert the board to an array of floats with 1.0 for the color
// to move, 0.0 for the opposite color and 0.5 for empty points
func (g *Grid) Neural(c Color) []float64 {
	rc := make([]float64, n*n)
	for xy := 0; xy < n*n; xy++ {
		switch g[xy] {
		case Empty:
			rc[xy] = 0.5
		case c:
			rc[xy] = 1.0
		default:
			rc[xy] = 0.0
		}
	}
	return rc
}

// rotate n*n plane clockwise
func Rotate(f []float64) []float64 {
	g := make([]float64, n*n)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			g[Xy(n-y-1, x)] = f[Xy(x, y)]
		}
	}

	return g
}

// flip n*n plane
func Flip(f []float64) []float64 {
	g := make([]float64, n*n)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			g[Xy(y, x)] = f[Xy(x, y)]
		}
	}

	return g
}
