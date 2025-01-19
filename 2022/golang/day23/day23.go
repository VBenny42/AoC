package day23

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type grid map[int]map[int]struct{}

type day23 struct {
	grid grid
}

func (g *grid) rounds(limit int) int {
	var (
		rounds             int
		neighbourIteration int
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

				d := neighbourIteration % 4
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

		neighbourIteration++

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

func (d *day23) Part1() int {
	return d.grid.rounds(10)
}

func (d *day23) Part2() int {
	return d.grid.rounds(10000000)
}

func Parse(filename string) *day23 {
	data := utils.ReadLines(filename)

	grid := make(grid, len(data))

	for y, line := range data {
		for x, char := range line {
			if char == '#' {
				grid.set(utils.Coord{X: x, Y: y})
			}
		}
	}

	return &day23{grid}
}

func Solve(filename string) {
	fmt.Println("ANSWER1: number of empty ground tiles after 10 rounds:", Parse(filename).Part1())
	fmt.Println("ANSWER2: round where no elf moves:", Parse(filename).Part2())
}
