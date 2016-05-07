package show

import (
	"neurogo/gogame"
	"fmt"
)

func ExampleEmptyBoard() {
	fmt.Println(show(&gogame.Grid{}))
	// Output:
	// ...
	// ...
	// ...
}

func ExampleBoardPosition() {
	g := gogame.Grid{}
	g.MakeMove(gogame.Xy(1, 1), gogame.Black).
		MakeMove(gogame.Xy(0, 0), gogame.White).
		MakeMove(gogame.Xy(1, 2), gogame.Black).
		MakeMove(gogame.Xy(0, 1), gogame.White)
	fmt.Println(show(&g))
	// Output:
	// O..
	// OX.
	// .X.
}
