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

func (g *grid) mkmove(xy int, c color) *grid {
	g[xy] = c
	return g
}
