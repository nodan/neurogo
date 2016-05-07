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
	moveType MoveType
	x, y     int
}

var emptyMove = Move{Pass, -1, -1}

type Position struct {
	turn  Color
	board Grid
	move  Move
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
	return g.CurrentPosition().board.Finished()
}

func (g *Game) Size() int {
	return Size
}

func (g *Game) Turn() Color {
	return g.CurrentPosition().turn
}

func (g *Game) Board() *Grid {
	return &g.CurrentPosition().board
}

func (g *Game) Move(x,y int) bool {
	player := g.Turn()
	cpos := g.CurrentPosition()
	grid := cpos.board
	if grid.MakeMove(Xy(x, y), player) == nil {
		return false
	}
	pos := Position{ Invert(player), grid, emptyMove }
	for _, p := range g.positions {
		if p.turn == pos.turn && reflect.DeepEqual(p.board, pos.board) {
			return false
		}
	}
	cpos.move = Move{Stone, x, y}
	g.positions = append(g.positions, pos)
	return true
}

func (g *Game) Pass() {
	p := g.CurrentPosition()
	p.move = Move{Pass, 0, 0}
	g.positions = append(g.positions, Position{ Invert(p.turn), p.board, emptyMove})
}
