package gogame

import (
	"reflect"
)

const (
	LegalMove = 0
	IllegalMove = 1
	SuperKo = 2
)

type IllegalMoveReason byte

type Position struct {
	turn  Color
	board Grid
}

type Game struct {
	positions []Position
}

func NewGame() *Game {
	ps := []Position{{Black, Grid{}}}
	return &Game{ps}
}

func (g *Game) currentPosition() *Position {
	return &g.positions[len(g.positions)-1]
}

func (g *Game) Finished() bool {
	return g.currentPosition().board.Finished()
}

func (g *Game) Size() int {
	return Size
}

func (g *Game) Turn() Color {
	return g.currentPosition().turn
}

func (g *Game) Move(x,y int) bool {
	player := g.Turn()
	grid := g.currentPosition().board
	if grid.MakeMove(Xy(x, y), player) == nil {
		return false
	}
	pos := Position{ Invert(player), grid }
	for _, p := range g.positions {
		if	reflect.DeepEqual(p, pos) {
			return false
		}
	}
	g.positions = append(g.positions, pos)
	return true
}
