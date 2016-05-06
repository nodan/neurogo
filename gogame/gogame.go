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

func (g *grid) mkmove(xy int, c color) *grid {
	g[xy] = c
	return g
}
