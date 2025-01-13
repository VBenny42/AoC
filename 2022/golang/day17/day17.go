package day17

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type day17 struct {
	heights    []int
	directions []utils.Coord
}

type rock struct {
	shapes    []utils.Coord
	leftMost  int
	rightMost int
	bottoms   []int
}

var rockShapes = []rock{
	minusShape,
	plusShape,
	lShape,
	lineShape,
	boxShape,
}

func (d *day17) simulate(numRocks int) int {
	var r rock
	rockIdx := 0
	dirIdx := 0
	for rockIdx < numRocks {
		fmt.Println("LOG: rockIdx:", rockIdx)
		r = rockShapes[rockIdx%5]
		for i := range r.shapes {
			r.shapes[i].Y += utils.MaxSlice(d.heights) + 3
		}

		settled := false
		for !settled {
			// 1. Handle jet movement
			direction := d.directions[dirIdx%len(d.directions)]
			dirIdx++

			// Try moving horizontally
			canMove := true
			newPositions := make([]utils.Coord, len(r.shapes))
			for i, shape := range r.shapes {
				newPos := shape
				if direction == utils.Left {
					newPos.X--
				} else {
					newPos.X++
				}
				// Check boundaries and collisions
				if newPos.X < 0 || newPos.X > 6 || d.checkCollision(newPos) {
					canMove = false
					break
				}
				newPositions[i] = newPos
			}

			if canMove {
				r.shapes = newPositions
			}

			// 2. Try falling
			canFall := true
			for _, bottomCell := range r.bottoms {
				pos := r.shapes[bottomCell]
				if pos.Y-1 <= d.heights[pos.X] {
					canFall = false
					break
				}
			}

			if !canFall {
				// Rock comes to rest
				for _, pos := range r.shapes {
					d.heights[pos.X] = max(d.heights[pos.X], pos.Y)
				}
				settled = true
			} else {
				// Move rock down
				for i := range r.shapes {
					r.shapes[i].Y--
				}
			}
		}
		rockIdx++
	}
	return utils.MaxSlice(d.heights)
}

func (d *day17) checkCollision(c utils.Coord) bool {
	return c.Y <= d.heights[c.X]
}

func (d *day17) Part1() int {
	return d.simulate(2022)
}

func Parse(filename string) *day17 {
	data := utils.ReadTrimmed(filename)

	directions := make([]utils.Coord, len(data))

	for i, c := range data {
		switch c {
		case '<':
			directions[i] = utils.Left
		case '>':
			directions[i] = utils.Right
		default:
			panic(fmt.Sprintf("Invalid character: %c", c))
		}
	}

	return &day17{
		directions: directions,
		heights:    []int{0, 0, 0, 0, 0, 0, 0},
	}
}

func Solve(filename string) {
	day := Parse(filename)
	day = Parse("../inputs/day17/sample-input.txt")

	fmt.Println("ANSWER1: highest column height after 2022 rocks:", day.Part1())
}
