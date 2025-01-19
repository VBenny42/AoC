package day23

import (
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type grid map[int]map[int]struct{}

func (g *grid) rounds(limit int) int {
	var (
		rounds  int
		dirIter int
	)

	for {
		rounds++

		var (
			proposed   = map[utils.Coord]utils.Coord{}
			cannotMove = map[utils.Coord]struct{}{}
		)

		var elves, still int

		for y, row := range *g {
			for x := range row {
				elves++
				current := utils.Coord{X: x, Y: y}

				neighbours := make([]bool, 12)

				// NW
				neighbours[0] = g.get(current.Add(utils.Up).Add(utils.Left))
				// N
				neighbours[1] = g.get(current.Add(utils.Up))
				// NE
				neighbours[2] = g.get(current.Add(utils.Up).Add(utils.Right))

				// SW
				neighbours[3] = g.get(current.Add(utils.Down).Add(utils.Left))
				// S
				neighbours[4] = g.get(current.Add(utils.Down))
				// SE
				neighbours[5] = g.get(current.Add(utils.Down).Add(utils.Right))

				// WS
				neighbours[6] = neighbours[3]
				// W
				neighbours[7] = g.get(current.Add(utils.Left))
				// WN
				neighbours[8] = neighbours[0]

				// ES
				neighbours[9] = neighbours[5]
				// E
				neighbours[10] = g.get(current.Add(utils.Right))
				// EN
				neighbours[11] = neighbours[2]

				allEmpty := true
				for _, neighbour := range neighbours {
					if neighbour {
						allEmpty = false
						break
					}
				}
				if allEmpty {
					still++
					continue
				}

				d := dirIter % 4
				for i := d; i < d+4; i++ {
					dir := i % 4

					allEmpty = true
					for _, neighbour := range neighbours[dir*3 : dir*3+3] {
						if neighbour {
							allEmpty = false
							break
						}
					}
					if allEmpty {
						next := current.Add(utils.Directions[dir])

						_, ok := proposed[next]

						if ok {
							// Elf already proposed to move here
							cannotMove[next] = struct{}{}
						} else {
							proposed[next] = current
						}

						break
					}
				}
			}
		}

		dirIter++

		for next, current := range proposed {
			_, ok := cannotMove[next]
			if ok {
				continue
			}

			g.delete(current)
			g.set(next)
		}

		if rounds == limit {
			var (
				bounds = g.bounds()
				height = bounds.maxY - bounds.minY + 1
				width  = bounds.maxX - bounds.minX + 1
				size   = height * width
			)

			return size - elves
		}

		if still == elves {
			return rounds
		}

	}
}

func (g *grid) set(c utils.Coord) {
	if _, ok := (*g)[c.Y]; !ok {
		(*g)[c.Y] = make(map[int]struct{})
	}
	(*g)[c.Y][c.X] = struct{}{}
}

func (g *grid) get(c utils.Coord) bool {
	if _, ok := (*g)[c.Y]; !ok {
		return false
	}
	_, ok := (*g)[c.Y][c.X]
	return ok
}

func (g *grid) delete(c utils.Coord) {
	if _, ok := (*g)[c.Y]; !ok {
		return
	}
	delete((*g)[c.Y], c.X)
}

type bounds struct {
	minX, maxX, minY, maxY int
}

func (g *grid) bounds() bounds {
	var minX, maxX, minY, maxY int

	for y, row := range *g {
		for x := range row {
			if y < minY {
				minY = y
			} else if y > maxY {
				maxY = y
			}

			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}
		}
	}

	return bounds{minX, maxX, minY, maxY}
}

func (g grid) String() string {
	var (
		res    []string
		bounds = g.bounds()
		width  = bounds.maxX - bounds.minX + 1
	)

	for i := bounds.minY; i <= bounds.maxY; i++ {
		res = append(res, strings.Repeat(".", width))
	}

	for y, row := range g {
		for x := range row {
			i := y - bounds.minY
			j := x - bounds.minX
			res[i] = res[i][:j] + "#" + res[i][j+1:]
		}
	}

	return strings.Join(res, "\n")
}
