package gogame

import (
	"bytes"
	//	"fmt"
)

type Group struct {
	side        color
	stoneCoords []byte
}

func leftOf(xy byte) (bool, byte) {
	if xy%n == 0 {
		return false, 0
	} else {
		return true, xy - 1
	}
}

func above(xy byte) (bool, byte) {
	if xy < n {
		return false, 0
	} else {
		return true, xy - n
	}
}

// findGroup finds the index of group of the expected side/color at the given coordinates.
// If it is not found, it returns -1, nil
func findGroup(grps []*Group, fn func(byte) (bool, byte), xy byte, side color) (bool, int, *Group) {
	ok, xy2 := fn(xy)
	if ok {
		for i, gp := range grps {
			r := bytes.IndexByte(gp.stoneCoords, xy2)
			if gp.side == side && r >= 0 {
				return true, i, gp
			}
		}
	}
	return false, -1, nil
}

func firstNonNil(gs ...*Group) *Group {
	for _, gp := range gs {
		if gp != nil {
			return gp
		}
	}
	return nil
}

// categorizeGroups finds all groups
func (g *grid) categorizeGroups() []*Group {
	grps := make([]*Group, 0)
	var i byte
	for i = 0; i < n*n; i++ {
		if g[i] == empty {
			continue
		}
		lok, _, lgp := findGroup(grps, leftOf, i, g[i])
		aok, aix, agp := findGroup(grps, above, i, g[i])
		if lok && aok {
			if lgp == agp {
				lgp.stoneCoords = append(lgp.stoneCoords, i)
			} else {
				lgp.stoneCoords = append(lgp.stoneCoords, agp.stoneCoords...)
				lgp.stoneCoords = append(lgp.stoneCoords, i)
				grps = append(grps[:aix], grps[aix+1:]...)
			}
		} else if lok || aok {
			gp := firstNonNil(lgp, agp)
			gp.stoneCoords = append(gp.stoneCoords, i)
		} else {
			grps = append(grps, &Group{g[i], []byte{i}})
		}
	}
	return grps
}
