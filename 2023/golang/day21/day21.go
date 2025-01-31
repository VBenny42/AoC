package day21

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	cell int
	grid [][]cell
)

type day21 struct {
	grid  grid
	start utils.Coord
}

const (
	plot cell = iota
	rock
)

func (g *grid) inBounds(c utils.Coord) bool {
	return c.X >= 0 && c.X < len((*g)[0]) && c.Y >= 0 && c.Y < len(*g)
}

func (g *grid) get(c utils.Coord) cell {
	return (*g)[c.Y][c.X]
}

func (d *day21) Part1And2() (part1, part2 int) {
	type node struct {
		coord     utils.Coord
		stepCount int
	}

	var (
		visited = make(map[utils.Coord]int)
		queue   = []node{{coord: d.start, stepCount: 0}}
	)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := visited[curr.coord]; ok {
			continue
		}
		visited[curr.coord] = curr.stepCount

		for _, dir := range utils.Directions {
			next := curr.coord.Add(dir)

			if !d.grid.inBounds(next) || d.grid.get(next) == rock {
				continue
			}

			_, ok := visited[next]
			if !ok {
				queue = append(queue, node{coord: next, stepCount: curr.stepCount + 1})
			}
		}
	}

	// Borrowed from
	// https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21
	var (
		allEvens, allOdds       int
		evenCorners, oddCorners int

		n = 202300

		even = n * n
		odd  = (n + 1) * (n + 1)
	)

	for _, v := range visited {
		if v%2 == 0 {
			allEvens++
			if v > 65 {
				evenCorners++
			} else {
				part1++
			}
		} else {
			allOdds++
			if v > 65 {
				oddCorners++
			}
		}
	}

	part2 = (odd * allOdds) + (even * allEvens) -
		((n + 1) * oddCorners) + (n * evenCorners)

	return
}

func Parse(filename string) *day21 {
	var (
		data  = utils.ReadLines(filename)
		grid  = make(grid, len(data))
		start utils.Coord
	)

	for y, line := range data {
		grid[y] = make([]cell, len(line))
		for x, c := range line {
			switch c {
			case '.':
				grid[y][x] = plot
			case '#':
				grid[y][x] = rock
			case 'S':
				grid[y][x] = plot
				start = utils.Coord{X: x, Y: y}
			}
		}
	}

	return &day21{grid: grid, start: start}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()
	fmt.Println("ANSWER1: number of plots that elf can reach in 64 steps:", part1)
	fmt.Println("ANSWER2: number of plots that elf can reach in 26501365 steps:", part2)
}
