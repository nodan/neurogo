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

const (
	Pass = 0
	Stone = 1
)

type MoveType byte

type Move struct {
	MoveType MoveType
	X, Y     int
}

var emptyMove = Move{Pass, -1, -1}

type Position struct {
	Turn  Color
	Board Grid
	Move  Move
}

type Game struct {
	positions []Position
}

func NewGame() *Game {
	ps := []Position{{Black, Grid{}, emptyMove}}
	return &Game{ps}
}

func (g *Game) Positions() []Position {
	return g.positions;
}

func (g *Game) CurrentPosition() *Position {
	return &g.positions[len(g.positions)-1]
}

func (g *Game) Finished() bool {
	return g.twoConsecutivePasses() || g.CurrentPosition().Board.Finished()
}

func (g *Game) twoConsecutivePasses() bool {
	n := len(g.positions)
	return n > 2 && g.positions[n-2].Move.MoveType == Pass && g.positions[n-3].Move.MoveType == Pass
}

func (g *Game) Size() int {
	return Size
}

func (g *Game) Turn() Color {
	return g.CurrentPosition().Turn
}

func (g *Game) Board() *Grid {
	return &g.CurrentPosition().Board
}

func (g *Game) Move(x,y int) bool {
	player := g.Turn()
	cpos := g.CurrentPosition()
	grid := cpos.Board
	if grid.MakeMove(Xy(x, y), player) == nil {
		return false
	}
	pos := Position{ Invert(player), grid, emptyMove }
	for _, p := range g.positions {
		if p.Turn == pos.Turn && reflect.DeepEqual(p.Board, pos.Board) {
			return false
		}
	}
	cpos.Move = Move{Stone, x, y}
	g.positions = append(g.positions, pos)
	return true
}

func (g *Game) Pass() {
	p := g.CurrentPosition()
	p.Move = Move{Pass, 0, 0}
	g.positions = append(g.positions, Position{ Invert(p.Turn), p.Board, emptyMove})
}
