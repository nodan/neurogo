package gogame

import (
	"testing"
	"reflect"
)

func TestCategorizeGroups(t *testing.T) {
	var g grid
	g.mkmove(xy(1, 1), white)
	result := g.categorizeGroups()
	if !reflect.DeepEqual(result, []Group{Group{white, []int{4}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

	g.mkmove(xy(1, 2), black)
	result = g.categorizeGroups()
	if !reflect.DeepEqual(result, []Group{Group{white, []int{4}},
		Group{black, []int{7}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

	g.mkmove(xy(1, 0), white).mkmove(xy(0, 2), black).mkmove(xy(2, 2), black)
	result = g.categorizeGroups()
	if !reflect.DeepEqual(result, []Group{Group{white, []int{1, 4}},
		Group{black, []int{6, 7, 8}}}) {
		t.Errorf("Got faulty answer %v\n", result)
	}

}
