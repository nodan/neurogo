package gogame

import (
	"testing"
	"reflect"
)

func TestCategorizeGroups(t *testing.T) {
	var g grid
	g.mkmove(xy(1, 1), white)
	result := g.categorizeGroups()
	if !reflect.DeepEqual(result, []*Group{&Group{white, []byte{4}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

	g.mkmove(xy(1, 2), black)
	result = g.categorizeGroups()
	if !reflect.DeepEqual(result, []*Group{&Group{white, []byte{4}},
		&Group{black, []byte{7}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

	g.mkmove(xy(1, 0), white).mkmove(xy(0, 2), black).mkmove(xy(2, 2), black)
	result = g.categorizeGroups()
	if !reflect.DeepEqual(result, []*Group{&Group{white, []byte{1, 4}},
		&Group{black, []byte{6, 7, 8}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

	g = grid{}
	g.mkmove(xy(0,0), white).mkmove(xy(0,1), white).mkmove(xy(0,2), white).
		mkmove(xy(2,0), white).mkmove(xy(2,1), white).mkmove(xy(2,2), white).
		mkmove(xy(1,2), white)
	result = g.categorizeGroups()
	if !reflect.DeepEqual(result, []*Group{&Group{white, []byte{0, 3, 6, 7, 2, 5, 8}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

}
