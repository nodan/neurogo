package gogame

import (
	"reflect"
	"testing"
)

func TestCategorizeGroups(t *testing.T) {
	var g Grid
	g.MakeMove(Xy(1, 1), White)
	result := g.categorizeGroups()
	if !reflect.DeepEqual(result, []*Group{{White, []byte{4}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

	g.MakeMove(Xy(1, 2), Black)
	result = g.categorizeGroups()
	if !reflect.DeepEqual(result, []*Group{{White, []byte{4}},
		{Black, []byte{7}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

	g.MakeMove(Xy(1, 0), White).MakeMove(Xy(0, 2), Black).MakeMove(Xy(2, 2), Black)
	result = g.categorizeGroups()
	if !reflect.DeepEqual(result, []*Group{{White, []byte{1, 4}},
		{Black, []byte{6, 7, 8}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

	g = Grid{}
	g.MakeMove(Xy(0, 0), White).MakeMove(Xy(0, 1), White).MakeMove(Xy(0, 2), White).
		MakeMove(Xy(2, 0), White).MakeMove(Xy(2, 1), White).MakeMove(Xy(2, 2), White).
		MakeMove(Xy(1, 2), White)
	result = g.categorizeGroups()
	if !reflect.DeepEqual(result, []*Group{{White, []byte{0, 3, 6, 7, 2, 5, 8}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

}
