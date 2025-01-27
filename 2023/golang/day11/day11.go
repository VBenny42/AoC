package day11

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
	"gonum.org/v1/gonum/stat/combin"
)

type (
	cell int
	grid [][]cell
)

type day11 struct {
	galaxies      []utils.Coord
	expansionRows []int
	expansionCols []int
}

const (
	empty cell = iota
	galaxy
)

func manhattanDistance(a, b utils.Coord) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func (d *day11) expandCoord(coord utils.Coord) func(int) utils.Coord {
	var row, col int

	for _, r := range d.expansionRows {
		if r < coord.Y {
			row++
		}
	}

	for _, c := range d.expansionCols {
		if c < coord.X {
			col++
		}
	}

	return func(factor int) utils.Coord {
		return utils.Crd(coord.X+(col*factor), coord.Y+(row*factor))
	}
}

func (d *day11) Part1And2() (part1, part2 int) {
	combinations := make([]int, 2)
	gen := combin.NewCombinationGenerator(len(d.galaxies), 2)

	for gen.Next() {
		gen.Combination(combinations)

		first, second := d.galaxies[combinations[0]], d.galaxies[combinations[1]]

		factor := 2 - 1

		var (
			firstExpanded  = d.expandCoord(first)
			secondExpanded = d.expandCoord(second)
		)

		part1 += manhattanDistance(
			firstExpanded(factor),
			secondExpanded(factor),
		)

		factor = 1000000 - 1

		part2 += manhattanDistance(
			firstExpanded(factor),
			secondExpanded(factor),
		)
	}

	return
}

func Parse(filename string) *day11 {
	data := utils.ReadLines(filename)
	galaxies := []utils.Coord{}

	grid := make(grid, len(data))
	for y, line := range data {
		grid[y] = make([]cell, len(line))
		for x, char := range line {
			if char == '#' {
				grid[y][x] = galaxy
				galaxies = append(galaxies, utils.Crd(x, y))
			}
		}
	}

	rows, cols := grid.expansionRowsAndCols()

	return &day11{
		galaxies:      galaxies,
		expansionRows: rows,
		expansionCols: cols,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println("ANSWER1: sum of shortest path between each pair of galaxies:", part1)
	fmt.Println("ANSWER2: sum of shortest path between each pair of galaxies with 1000000 expansions:", part2)
}
