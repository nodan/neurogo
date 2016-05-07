package gogame

import (
	"fmt"
)

func ExampleEmptyBoard() {
	fmt.Println((&Grid{}).Show())
	// Output:
	// ...
	// ...
	// ...
}

func ExampleBoardPosition() {
	g := Grid{}
	g.MakeMove(Xy(1, 1), Black).
		MakeMove(Xy(0, 0), White).
		MakeMove(Xy(1, 2), Black).
		MakeMove(Xy(0, 1), White)
	fmt.Println(g.Show())
	// Output:
	// O..
	// OX.
	// .X.
}
