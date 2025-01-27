package day11

func (g *grid) expansionRowsAndCols() (rows, cols []int) {
	for y := 0; y < len(*g); y++ {
		shouldExpand := true
		for x := 0; x < len((*g)[0]); x++ {
			if (*g)[y][x] == galaxy {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			rows = append(rows, y)
		}
	}

	for x := 0; x < len((*g)[0]); x++ {
		shouldExpand := true
		for y := 0; y < len(*g); y++ {
			if (*g)[y][x] == galaxy {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			cols = append(cols, x)
		}
	}

	return
}
