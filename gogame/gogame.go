package gogame

import "os"
import "testing"

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

func mkmove(g *grid, xy int, c color) {
	g[xy] = c
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
