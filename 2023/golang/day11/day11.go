package day11

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	cell int
	grid [][]cell
)

type day11 struct {
	grid     grid
	galaxies []utils.Coord
}

const (
	empty cell = iota
	galaxy
	expansion
)

func (d *day11) Part1And2() (part1, part2 int) {
	var (
		memo = map[utils.Coord]map[utils.Coord]stepWithExpansion{}
		seen = map[[2]utils.Coord]bool{}
	)

	for _, galaxy := range d.galaxies {
		memo[galaxy] = map[utils.Coord]stepWithExpansion{}
	}

	for _, start := range d.galaxies {
		for _, end := range d.galaxies {
			if start == end {
				continue
			}

			first, second := start, end
			if start.X > end.X || (start.X == end.X && start.Y > end.Y) {
				first, second = end, start
			}

			if seen[[2]utils.Coord{first, second}] {
				continue
			}
			seen[[2]utils.Coord{first, second}] = true

			steps, err := d.grid.bfs(first, second, memo)
			if err != nil {
				fmt.Println(err)
				return
			}

			part1 += steps.actual + (steps.expansions * 2)
			part2 += steps.actual + (steps.expansions * 1000000)
		}
	}
	return
}

func Parse(filename string) *day11 {
	data := utils.ReadLines(filename)

	grid := make(grid, len(data))
	for i, line := range data {
		grid[i] = make([]cell, len(line))
		for j, char := range line {
			if char == '#' {
				grid[i][j] = galaxy
			} else {
				grid[i][j] = empty
			}
		}
	}

	expandedGrid := grid.expand()
	galaxies := []utils.Coord{}

	for y, row := range grid {
		for x, cell := range row {
			if cell == galaxy {
				galaxies = append(galaxies, utils.Crd(x, y))
			}
		}
	}

	return &day11{
		grid:     expandedGrid,
		galaxies: galaxies,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println("ANSWER1: sum of shortest path between each pair of galaxies:", part1)
	fmt.Println("ANSWER2: sum of shortest path between each pair of galaxies with 1000000 expansions:", part2)
}
