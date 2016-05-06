package gogame

const (
	// board size
	n = 3

	// colors
	empty = 0
	black = 1
	white = 2
)

type color byte

// the board
type grid [n * n]color

// transform a board coordinate into a linear array index
func xy(x, y int) int {
	return n*y + x
}

// invert a color black<->white
func invert(c color) color {
	switch c {
	case black:
		return white
	case white:
		return black
	default:
		return empty
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
func (g *grid) findLiberties(xy int, max int) int {
	libs := 0
	c := g[xy]
	opposite := invert(c)
	g[xy] = opposite // don't look here again

	// look at the neighbors
	for _, nxy := range neighbors(xy) {
		switch g[nxy] {
		case empty:
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
func (g *grid) liberties(xy int, max int) int {
	t := *g
	return t.findLiberties(xy, max)
}

// remove a chain of stones
func (g *grid) remove(xy int) *grid {
	c := g[xy]
	g[xy] = empty
	for _, nxy := range neighbors(xy) {
		if g[nxy] == c {
			g.remove(nxy)
		}
	}

	return g
}

// count liberties of a chain of stones
func (counter *grid) count(xy int, g *grid) *grid {
	if c := g[xy]; c != empty {
		g[xy] = empty
		counter[xy]++
		for _, nxy := range neighbors(xy) {
			if g[nxy] == c {
				counter.count(nxy, g)
			}
		}
	}
	return counter
}

// play a move
func (g *grid) mkmove(xy int, c color) *grid {
	if g[xy] != empty {
		// don't play on non-empty points
		return nil
	}

	// play a move
	g[xy] = c

	// check neighbors
	for _, nxy := range neighbors(xy) {
		t := *g
		if t[nxy] == invert(c) && t.findLiberties(nxy, 1) == 0 {
			// remove captured stones
			g.remove(nxy)
		}
	}

	// check liberties of the move played
	t := *g
	if t.findLiberties(xy, 1) == 0 {
		// undo the move
		g[xy] = empty

		// illegal move, no liberties
		return nil
	}

	return g
}

// check if the game is finished in the sense of there not being to adjacent empty points and
// every group having exactly two liberties
func (g *grid) finished() bool {
	var c grid
	for xy := 0; xy < n*n; xy++ {
		// find empty points
		if g[xy] == empty {
			t := *g
			nl := neighbors(xy)
			for _, nxy := range nl {
				// same color for all adjacent points
				if g[nxy] == empty || g[nxy] != g[nl[0]] {
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

	// find non-empty points
	for xy := 0; xy < n*n; xy++ {
		// all chains must have 2 liberties
		if g[xy] != empty && c[xy] != 2 {
			return false
		}
	}

	return true
}

func (g *grid) legal() bool {
	for xy := 0; xy < n*n; xy++ {
		println(xy, g[xy], g.liberties(xy, 1))
		if g[xy] != empty && g.liberties(xy, 1) == 0 {
			return false
		}
	}
	return true
}
