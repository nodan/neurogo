package gogame

import (
	"testing"
//	"fmt"
)

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

	// Test
	// OOO
	// OOO
	// .O.
	// is finished
	g = grid{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 2; y++ {
			g.mkmove(xy(x, y), white)
		}
	}
	g.mkmove(xy(1, 2), white)
	if !g.finished() {
		t.Errorf("Game not finished")
	}

	// Test
	// OOO
	// OOO
	// OO.
	// is not finished
	g = grid{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			g.mkmove(xy(x, y), white)
		}
	}
	g.mkmove(xy(2, 2), empty)
	if g.finished() {
		t.Errorf("Game not finished")
	}

	g = grid{}
	if !g.legal() {
		t.Errorf("Empty board not legal")
	}

	g = grid{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			g.mkmove(xy(x, y), white)
		}
	}
	if !g.legal() {
		t.Errorf("Almost full board not legal")
	}

	g[8] = white
	if g.legal() {
		t.Errorf("Full board legal")
	}

	g = grid{}
	g.mkmove(xy(1,0), white).mkmove(xy(0,1), white)
	g[0] = black
	if g.legal() {
		t.Errorf("Captured stone not recognized")
	}

}
