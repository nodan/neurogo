package gogame

import "testing"

func TestXyAndMkMove(t *testing.T) {
	var g grid
	mkmove(&g, xy(1, 1), white)
	expectedGrid := grid{0, 0, 0, 0, 2, 0, 0, 0, 0}
	if g != expectedGrid {
		t.Errorf("Expected %v, but got %v\n", expectedGrid, g)
	}
}
