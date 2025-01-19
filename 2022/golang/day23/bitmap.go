package day23

import (
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

const offset = 20

type bitmap [160][160]bool

func (b *bitmap) rounds(limit int) int {
	var (
		rounds    int
		direction int
		current   utils.Coord
	)

	var (
		neighbours = make([]bool, 12)
		proposed   = map[utils.Coord]utils.Coord{}
		cannotMove = map[utils.Coord]struct{}{}
	)

	var elves, still int

	for {
		rounds++

		clear(proposed)
		clear(cannotMove)

		elves, still = 0, 0

		for y, row := range *b {
			for x := range row {
				if !(*b)[y][x] {
					continue
				}
				elves++
				current = utils.Coord{X: x, Y: y}

				// NW
				neighbours[0] = b.get(current.Add(utils.Up).Add(utils.Left))
				// N
				neighbours[1] = b.get(current.Add(utils.Up))
				// NE
				neighbours[2] = b.get(current.Add(utils.Up).Add(utils.Right))

				// SW
				neighbours[3] = b.get(current.Add(utils.Down).Add(utils.Left))
				// S
				neighbours[4] = b.get(current.Add(utils.Down))
				// SE
				neighbours[5] = b.get(current.Add(utils.Down).Add(utils.Right))

				// WS
				neighbours[6] = neighbours[3]
				// W
				neighbours[7] = b.get(current.Add(utils.Left))
				// WN
				neighbours[8] = neighbours[0]

				// ES
				neighbours[9] = neighbours[5]
				// E
				neighbours[10] = b.get(current.Add(utils.Right))
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

				for i := direction; i < direction+4; i++ {
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

		if still == elves {
			return rounds
		}

		direction = (direction + 1) % 4

		for next, current := range proposed {
			_, ok := cannotMove[next]
			if ok {
				continue
			}

			b.unset(current)
			b.set(next)
		}

		if rounds == limit {
			var (
				bounds = b.bounds()
				height = bounds.maxY - bounds.minY + 1
				width  = bounds.maxX - bounds.minX + 1
				size   = height * width
			)

			return size - elves
		}

	}
}

func (b *bitmap) set(c utils.Coord) {
	(*b)[c.Y][c.X] = true
}

func (b *bitmap) get(c utils.Coord) bool {
	return (*b)[c.Y][c.X]
}

func (b *bitmap) unset(c utils.Coord) {
	(*b)[c.Y][c.X] = false
}

func (b *bitmap) bounds() bounds {
	var minX, maxX, minY, maxY int
	minX, minY = offset, offset

	for y, row := range *b {
		for x, cell := range row {
			if cell {
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
	}

	return bounds{minX - offset, maxX - offset, minY - offset, maxY - offset}
}

func (b bitmap) String() string {
	var builder strings.Builder

	for y := range b {
		for x := range b[y] {
			if b[y][x] {
				builder.WriteByte('#')
			} else {
				builder.WriteByte('.')
			}
		}
		builder.WriteByte('\n')
	}

	return builder.String()
}
