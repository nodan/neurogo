package gogame

import (
	"bytes"
	"fmt"
	"strings"
)

func (g *Grid) Show() string {
	var b bytes.Buffer
	var m = map[Color]byte{
		Empty: '.',
		White: 'O',
		Black: 'X',
	}
	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			b.WriteByte(m[g[Xy(x, y)]])
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (p *Position) ShowPosition() string {
	var m = map[Color]string{
		White: "White",
		Black: "Black",
	}
	return m[p.Turn] + " to play\n" + p.Board.Show()
}

func (g *Game) ShowCurrentPosition() string {
	return g.CurrentPosition().ShowPosition()
}

func (g *Game) ShowGame() string {
	result := make([]string, len(g.positions))
	for _, p := range g.positions {
		var mv string
		if p.Move.MoveType == Pass {
			mv = "passed\n"
		} else {
			mv = fmt.Sprintf("@%v:%v\n", p.Move.X, p.Move.Y)
		}
		result = append(result, p.ShowPosition()+mv)
	}
	return strings.Join(result, "\n")
}
