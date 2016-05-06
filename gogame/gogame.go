package gogame

const (
	n = 3
	empty = 0
	black = 1
	white = 2
)

type color byte
type grid [n*n] color

func xy(x, y int) int {
	return n*y+x
}

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

func neighbors(xy int) []int {
	rc := make([]int, 0, 4)
	if xy%n>=1 {
		rc = append(rc, xy-1)
	}

	if xy%n+1!=n {
		rc = append(rc, xy+1)
	}

	if xy>=n {
		rc = append(rc, xy-n)
	}

	if xy+n<n*n {
		rc = append(rc, xy+n)
	}

	return rc
}

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
			if libs>=max {
				return libs
			}
		case c:
			// recursively count liberties of neighbor
			libs += g.findLiberties(nxy, max-libs)

			// count up to max libs
			if libs>=max {
				return libs
			}
		}
	}

	// the liberties found so far
	return libs
}

func (g *grid) liberties(xy int, max int) int {
	t := *g
	return t.findLiberties(xy, max)
}

func (g *grid) mkmove(xy int, c color) *grid {
	g[xy] = c
	t := *g
	for _, nxy := range neighbors(xy) {
		if t[nxy]==invert(c) && t.findLiberties(nxy, 1)==0 {
			// remove captured stones
			g[nxy] = empty
		}
	}

	if t.findLiberties(xy, 1)==0 {
		// illegal move, no liberties
		g[xy] = empty
	}

	return g
}
