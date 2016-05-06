package gogame

import "testing"

func TestXyAndMkMove(t *testing.T) {
	if invert(black)!=white || invert(white)!=black || invert(empty)!=empty {
		t.Errorf("Failed to invert colors\n")
	}

	var g grid
	g.mkmove(xy(1, 1), white)
	expectedGrid := grid{empty, empty, empty, empty, white, empty, empty, empty, empty}
	if g != expectedGrid {
		t.Errorf("Expected %v, but got %v\n", expectedGrid, g)
	}
}
