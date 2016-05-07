package gogame

import (
	"testing"
	//	"fmt"
)

func TestXyAndMakeMove(t *testing.T) {
	if Invert(Black) != White || Invert(White) != Black || Invert(Empty) != Empty {
		t.Errorf("Failed to invert colors\n")
	}

	var g Grid
	g.MakeMove(xy(1, 1), White)
	expectedGrid := Grid{Empty, Empty, Empty, Empty, White, Empty, Empty, Empty, Empty}
	if g != expectedGrid {
		t.Errorf("Expected %v, but got %v\n", expectedGrid, g)
	}

	g.MakeMove(xy(0, 1), Black)
	g.MakeMove(xy(1, 0), Black)
	expectedGrid = Grid{Empty, Black, Empty, Black, White, Empty, Empty, Empty, Empty}
	if g != expectedGrid {
		t.Errorf("Expected %v, but got %v\n", expectedGrid, g)
	}

	if l := g.liberties(xy(1, 1), 4); l != 2 {
		t.Errorf("Expected 2 liberties, but got %d\n", l)
	}

	g.MakeMove(xy(2, 1), Black)
	g.MakeMove(xy(1, 2), Black)
	expectedGrid = Grid{Empty, Black, Empty, Black, Empty, Black, Empty, Black, Empty}
	if g != expectedGrid {
		t.Errorf("Expected %v, but got %v\n", expectedGrid, g)
	}

	if g.MakeMove(xy(0, 0), White) != nil {
		t.Errorf("Allowed illegal move at (0, 0)")
	}

	if g.MakeMove(xy(1, 1), White) != nil {
		t.Errorf("Allowed illegal move at (1, 1)")
	}

	if g.MakeMove(xy(2, 2), White) != nil {
		t.Errorf("Allowed illegal move at (2, 2)")
	}

	if g.Finished() {
		t.Errorf("Game finished")
	}

	g.MakeMove(xy(2, 0), Black)
	g.MakeMove(xy(1, 1), Black)
	g.MakeMove(xy(0, 2), Black)

	if !g.Finished() {
		t.Errorf("Game not finished")
	}

	// Test
	// OOO
	// OOO
	// .O.
	// is finished
	g = Grid{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 2; y++ {
			g.MakeMove(xy(x, y), White)
		}
	}
	g.MakeMove(xy(1, 2), White)
	if !g.Finished() {
		t.Errorf("Game not finished")
	}

	// Test
	// OOO
	// OOO
	// OO.
	// is not finished
	g = Grid{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			g.MakeMove(xy(x, y), White)
		}
	}
	g.MakeMove(xy(2, 2), Empty)
	if g.Finished() {
		t.Errorf("Game not finished")
	}

	g = Grid{}
	if !g.Legal() {
		t.Errorf("Empty board not legal")
	}

	g = Grid{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			g.MakeMove(xy(x, y), White)
		}
	}
	if !g.Legal() {
		t.Errorf("Almost full board not legal")
	}

	g[8] = White
	if g.Legal() {
		t.Errorf("Full board legal")
	}

	g = Grid{}
	g.MakeMove(xy(1, 0), White).MakeMove(xy(0, 1), White)
	g[0] = Black
	if g.Legal() {
		t.Errorf("Captured stone not recognized")
	}

}
