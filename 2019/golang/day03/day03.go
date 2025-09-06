package day03

import (
	"fmt"
	"image"
	"strings"

	"github.com/VBenny42/AoC/2019/golang/utils"
)

type day03 struct {
	path1, path2 []string
}

func (d *day03) Part1And2() (minDist int, minEffort int) {
	var (
		// Track which wires have visited each point
		wire1Points = make(map[image.Point]bool)
		grid1       = make(map[image.Point]int)
		grid2       = make(map[image.Point]int)
		// Trace first wire
		previous = image.Point{0, 0}
		effort   int
	)

	for _, move := range d.path1 {
		var (
			dir   = move[0]
			dist  = utils.Atoi(move[1:])
			delta image.Point
		)

		switch dir {
		case 'U':
			delta = utils.Up
		case 'D':
			delta = utils.Down
		case 'L':
			delta = utils.Left
		case 'R':
			delta = utils.Right
		}
		for i := 0; i < dist; i++ {
			previous = previous.Add(delta)
			wire1Points[previous] = true
			effort++
			if _, exists := grid1[previous]; !exists {
				grid1[previous] = effort
			}
		}
	}

	previous = image.Point{0, 0}
	effort = 0
	minDist = -1
	minEffort = -1

	for _, move := range d.path2 {
		var (
			dir   = move[0]
			dist  = utils.Atoi(move[1:])
			delta image.Point
		)

		switch dir {
		case 'U':
			delta = utils.Up
		case 'D':
			delta = utils.Down
		case 'L':
			delta = utils.Left
		case 'R':
			delta = utils.Right
		}
		for i := 0; i < dist; i++ {
			previous = previous.Add(delta)
			effort++
			if _, exists := grid2[previous]; !exists {
				grid2[previous] = effort
			}
			// Check if this point was visited by wire1
			if wire1Points[previous] {
				dist := utils.Abs(previous.X) + utils.Abs(previous.Y)
				if minDist == -1 || dist < minDist {
					minDist = dist
				}
				totalEffort := grid1[previous] + grid2[previous]
				if minEffort == -1 || totalEffort < minEffort {
					minEffort = totalEffort
				}
			}
		}
	}

	return minDist, minEffort
}

func Parse(filename string) *day03 {
	lines := utils.ReadLines(filename)

	return &day03{
		path1: strings.Split(lines[0], ","),
		path2: strings.Split(lines[1], ","),
	}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println(
		"ANSWER1: Manhattan distance from central port to closest intersection:",
		part1,
	)
	fmt.Println(
		"ANSWER2: Fewest combined steps the wires must take to reach an intersection:",
		part2,
	)
}
