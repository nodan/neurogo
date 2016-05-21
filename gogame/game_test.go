package gogame

import (
	"fmt"
	"testing"
)

func TestNewGameNotFinished(t *testing.T) {
	if NewGame().Finished() {
		t.Errorf("Empty board is finished")
	}
}

func TestLegalMoves(t *testing.T) {
	g := NewGame()
	if !g.Move(1, 1) {
		t.Errorf("Black @ 1,1 illegal")
	}
	if !g.Move(0, 1) {
		t.Errorf("White @ 0,1 illegal")
	}
	if g.Move(1, 1) {
		t.Errorf("Black on black @ 1,1 legal")
	}
	if !g.Move(2, 1) {
		t.Errorf("Black @ 2,1 illlegal")
	}
	if !g.Move(1, 0) {
		t.Errorf("White @ 1,0 illegal")
	}
	if g.Move(0, 0) {
		t.Errorf("Black suicide @ 0,0 legal")
	}
	if len(g.positions) != 5 {
		t.Errorf("5 positions expected but got %d", len(g.positions))
	}
}

func TestSuperKo(t *testing.T) {
	g := NewGame()
	g.Move(1, 1)
	g.Move(1, 0)
	g.Move(0, 2)
	g.Move(0, 1)
	g.Move(0, 0)
	if g.Move(0, 1) {
		t.Errorf("Ko not recognized")
	}
	println(g.ShowGame())
}

func TestPass(t *testing.T) {
	g := NewGame()
	g.Pass()
	if g.Turn() != White {
		t.Errorf("Not white's turn, but %v", g.Turn())
	}
}

func ExampleBlackMovesTwice() {
	g := NewGame()
	g.Move(1, 1)
	g.Pass()
	g.Move(2, 1)
	fmt.Println(g.ShowGame())
	// Output:
	// Black to play
	// ...
	// ...
	// ...
	// @1:1
	//
	// White to play
	// ...
	// .X.
	// ...
	// passed
	//
	// Black to play
	// ...
	// .X.
	// ...
	// @2:1
	//
	// White to play
	// ...
	// .XX
	// ...
	// passed
}
