package gogame

type Group struct {
	side   color
	coords []int
}

func leftOf(coords int) (bool, int) {
	if coords%n == 0 {
		return false, -1
	} else {
		return true, coords - 1
	}
}

func above(coords int) (bool, int) {
	if coords < n {
		return false, -1
	} else {
		return true, coords - n
	}
}

// categorizeGroups finds all groups
func (g *grid) categorizeGroups() []Group {
	grps := make([]Group, 0)
	for i := 0; i < n*n; i++ {
		if g[i] == empty {
			continue
		}
		grps = append(grps, Group{g[i], []int{i}})
	}
	return grps
}
