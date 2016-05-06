package gogame

import (
	"bytes"
	"fmt"
)


type Group struct {
	side   color
	stoneCoords []byte
}

func leftOf(coords byte) (bool, byte) {
	if coords%n == 0 {
		return false, 0
	} else {
		return true, coords - 1
	}
}

func above(coords byte) (bool, byte) {
	println("above", coords)
	if coords < n {
		println("nothing above")
		return false, 0
	} else {
		println("above returns", coords - n)
		return true, coords - n
	}
}

// findGroup finds the index of group of the expected side/color at the given coordinates.
// If it is not found, it returns -1, nil
func findGroup(grps []Group, ok bool, coords byte, side color) (bool, int, *Group) {
	fmt.Printf("findGroup %v %v %v\n", grps, coords, side)
	if ok {
		for i, gp := range grps {
			r := bytes.IndexByte(gp.stoneCoords, coords)
			if gp.side == side && r >= 0 {
				println("found", i)
				return true, i, &gp
			}
		}
	}
	println("nothing found")
	return false, -1, nil
}



// categorizeGroups finds all groups
func (g *grid) categorizeGroups() []Group {
	grps := make([]Group, 0)
	var i byte
	newGrp := func() {
		grps = append(grps, Group{g[i], []byte{i}})
	}
	for i = 0; i < n*n; i++ {
		if g[i] == empty {
			println("skip", i)
			continue
		}
		println("@", i)
		var ok bool
		var coords byte
		ok, coords = leftOf(i)
		lok, _, lgp := findGroup(grps, ok, coords, g[i])
		ok, coords = above(i)
		aok, _, agp := findGroup(grps, ok, coords, g[i])
		if lok || aok {
			var gp *Group
			if lok {
				gp = lgp
			} else {
				gp = agp
			}
			println(gps)
			fmt.Printf("GP: %v\n", gp)
			println(gp)
			gp.stoneCoords = append(gp.stoneCoords, i)
			fmt.Printf("GP: %v\n", gp)
			println(gp)
			fmt.Printf("GPS: %v\n", grps)
			println(gps)
		} else {
			newGrp()
		}
	}
	return grps
}
