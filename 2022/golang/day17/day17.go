package day17

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type day17 struct {
	directions []utils.Coord
}

type rock []utils.Coord

const width = 7

var rockShapes = []rock{
	minusShape,
	plusShape,
	lShape,
	lineShape,
	boxShape,
}

type seenKey struct {
	rockIndex      int
	directionIndex int
}

type seenValue struct {
	seenKeyCount  int
	rockCount     int
	highestColumn int
}

type state struct {
	directionCount int
	rockCount      int
	highestColumn  int
	grid           [][width]bool
	currentCoord   utils.Coord
	addedByRepeats int
	seen           map[seenKey]seenValue
}

func (s *state) isValid(newCoord utils.Coord, rock rock) bool {
	for i := range rock {
		x := newCoord.X + rock[i].X
		y := newCoord.Y + rock[i].Y
		for len(s.grid) <= y {
			s.grid = append(s.grid, [width]bool{})
		}
		if x >= width || s.grid[y][x] {
			return false
		}
	}

	return true
}

func (s *state) simulate(rockNum int, directions []utils.Coord) {
	for s.rockCount < rockNum {
		rock := rockShapes[s.rockCount%len(rockShapes)]
		s.currentCoord = utils.Coord{X: 2, Y: s.highestColumn + 3}

		for {
			direction := directions[s.directionCount%len(directions)]
			newCoord := utils.Coord{X: s.currentCoord.X, Y: s.currentCoord.Y}

			switch direction {
			case utils.Left:
				newCoord.X--
				if newCoord.X < 0 {
					newCoord.X = 0
				}
			case utils.Right:
				newCoord.X++
				if newCoord.X >= width {
					newCoord.X = width - 1
				}
			}

			if s.isValid(newCoord, rock) {
				s.currentCoord = newCoord
			}
			s.directionCount++

			newCoord = utils.Coord{X: s.currentCoord.X, Y: s.currentCoord.Y - 1}
			if s.currentCoord.Y == 0 || !s.isValid(newCoord, rock) {
				break
			}

			s.currentCoord = newCoord
		}

		for i := range rock {
			x := s.currentCoord.X + rock[i].X
			y := s.currentCoord.Y + rock[i].Y

			for len(s.grid) <= y {
				s.grid = append(s.grid, [width]bool{})
			}

			s.grid[y][x] = true
			s.highestColumn = max(s.highestColumn, y+1)
		}

		if s.addedByRepeats == 0 {
			key := seenKey{
				rockIndex:      s.rockCount % len(rockShapes),
				directionIndex: s.directionCount % len(directions),
			}

			if val, ok := s.seen[key]; ok {
				if val.seenKeyCount == 2 {
					highestColumnDiff := s.highestColumn - val.highestColumn
					rockCountDiff := s.rockCount - val.rockCount
					repeats := (rockNum - s.rockCount) / rockCountDiff
					s.addedByRepeats = highestColumnDiff * repeats
					s.rockCount += rockCountDiff * repeats
				}
				val.seenKeyCount++
				val.rockCount = s.rockCount
				val.highestColumn = s.highestColumn
				s.seen[key] = val
			} else {
				s.seen[key] = seenValue{
					seenKeyCount:  1,
					rockCount:     s.rockCount,
					highestColumn: s.highestColumn,
				}
			}
		}

		s.rockCount++
	}
}

func (d *day17) Part1() int {
	s := state{}
	s.seen = make(map[seenKey]seenValue)

	s.simulate(2022, d.directions)

	return s.highestColumn + s.addedByRepeats
}

func (d *day17) Part2() int {
	s := state{}
	s.seen = make(map[seenKey]seenValue)

	s.simulate(1000000000000, d.directions)

	return s.highestColumn + s.addedByRepeats
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

	return &day17{directions}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: highest column height after 2022 rocks:", day.Part1())
	fmt.Println("ANSWER2: highest column height after 1_000_000_000_000 rocks:", day.Part2())
}
