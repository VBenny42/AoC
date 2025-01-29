package day14

import "github.com/VBenny42/AoC/2023/golang/utils"

func (g *grid) moveRockMaxUp(rock utils.Coord) {
	currY := rock.Y
	for currY > 0 {
		// look up, if can't move up, return current position
		if (*g)[currY-1][rock.X] != empty {
			return
		}
		(*g)[currY][rock.X] = empty
		(*g)[currY-1][rock.X] = rounded
		currY--
	}
}

func (g *grid) moveAllRocksUp() {
	for y := 0; y < len((*g)); y++ {
		for x := 0; x < len((*g)[0]); x++ {
			if (*g)[y][x] == rounded {
				g.moveRockMaxUp(utils.Crd(x, y))
			}
		}
	}
}

func (g *grid) calculateNorthLoad() (load int) {
	for x := 0; x < len((*g)[0]); x++ {
		for y := 0; y < len((*g)); y++ {
			if (*g)[y][x] == rounded {
				load += len((*g)) - y
			}
		}
	}

	return
}

func (g *grid) rotateClockWise() {
	n := len(*g)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := (*g)[i][j]
			(*g)[i][j] = (*g)[n-1-j][i]
			(*g)[n-1-j][i] = (*g)[n-1-i][n-1-j]
			(*g)[n-1-i][n-1-j] = (*g)[j][n-1-i]
			(*g)[j][n-1-i] = temp
		}
	}
}

func (g *grid) hash() uint32 {
	hash := uint32(1)
	for _, line := range *g {
		for _, c := range line {
			hash = hash*31 + uint32(c)
		}
	}
	return hash
}
