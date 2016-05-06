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

	g.mkmove(xy(0, 1), black)
	g.mkmove(xy(1, 0), black)
	expectedGrid = grid{empty, black, empty, black, white, empty, empty, empty, empty}
	if g != expectedGrid {
		t.Errorf("Expected %v, but got %v\n", expectedGrid, g)
	}

	if l := g.liberties(xy(1,1), 4); l!=2 {
		t.Errorf("Expected 2 liberties, but got %d\n", l)
	}

	g.mkmove(xy(2, 1), black)
	g.mkmove(xy(1, 2), black)
	expectedGrid = grid{empty, black, empty, black, empty, black, empty, black, empty}
	if g != expectedGrid {
		t.Errorf("Expected %v, but got %v\n", expectedGrid, g)
	}

	if g.mkmove(xy(0, 0), white)!=nil {
		t.Errorf("Allowed illegal move at (0, 0)")
	}

	if g.mkmove(xy(1, 1), white)!=nil {
		t.Errorf("Allowed illegal move at (1, 1)")
	}

	if g.mkmove(xy(2, 2), white)!=nil {
		t.Errorf("Allowed illegal move at (2, 2)")
	}

	if g.finished() {
		t.Errorf("Game finished")
	}

	g.mkmove(xy(2, 0), black)
	g.mkmove(xy(1, 1), black)
	g.mkmove(xy(0, 2), black)

	if !g.finished() {
		t.Errorf("Game not finished")
	}
}
