package show

import (
	"neurogo/gogame"
	"bytes"
)

func show(g *gogame.Grid) string {
	var b bytes.Buffer
	var m = map[gogame.Color]byte{
		gogame.Empty: '.',
		gogame.White: 'O',
		gogame.Black: 'X',
	}
	for y := 0; y < gogame.Size; y++ {
		for x := 0; x < gogame.Size; x++ {
			b.WriteByte(m[g[gogame.Xy(x, y)]])
		}
		b.WriteByte('\n')
	}
	return b.String()
}
