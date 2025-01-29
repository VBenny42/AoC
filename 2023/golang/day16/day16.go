package day16

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	cell int
	grid [][]cell
)

type day16 struct {
	grid grid
}

const (
	empty cell = iota
	vSplit
	hSplit
	rightMirror
	leftMirror
)

var mapRune = map[rune]cell{
	'.':  empty,
	'|':  vSplit,
	'-':  hSplit,
	'/':  rightMirror,
	'\\': leftMirror,
}

func (g *grid) inBounds(c utils.Coord) bool {
	return c.Y >= 0 && c.Y < len(*g) && c.X >= 0 && c.X < len((*g)[0])
}

type state struct {
	pos utils.Coord
	dir utils.Coord
}

func (d *day16) simulateBeams(start state, visited map[state]struct{}, energized map[utils.Coord]struct{}) int {
	clear(visited)
	clear(energized)
	queue := []state{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := visited[curr]; ok {
			continue
		}
		visited[curr] = struct{}{}

		energized[curr.pos] = struct{}{}

		// Determine the next directions based on the current cell and direction
		cellType := d.grid[curr.pos.Y][curr.pos.X]

		var nextDirs []utils.Coord

		switch {
		case cellType&(1<<empty) != 0:
			nextDirs = []utils.Coord{curr.dir}
		case cellType&(1<<vSplit) != 0:
			if curr.dir == utils.Up || curr.dir == utils.Down {
				nextDirs = append(nextDirs, curr.dir)
			} else {
				nextDirs = append(nextDirs, utils.Up, utils.Down)
			}
		case cellType&(1<<hSplit) != 0:
			if curr.dir == utils.Left || curr.dir == utils.Right {
				nextDirs = append(nextDirs, curr.dir)
			} else {
				nextDirs = append(nextDirs, utils.Left, utils.Right)
			}
		case cellType&(1<<rightMirror) != 0:
			switch curr.dir {
			case utils.Right:
				nextDirs = append(nextDirs, utils.Up)
			case utils.Left:
				nextDirs = append(nextDirs, utils.Down)
			case utils.Up:
				nextDirs = append(nextDirs, utils.Right)
			case utils.Down:
				nextDirs = append(nextDirs, utils.Left)
			}
		case cellType&(1<<leftMirror) != 0:
			switch curr.dir {
			case utils.Right:
				nextDirs = append(nextDirs, utils.Down)
			case utils.Left:
				nextDirs = append(nextDirs, utils.Up)
			case utils.Up:
				nextDirs = append(nextDirs, utils.Left)
			case utils.Down:
				nextDirs = append(nextDirs, utils.Right)
			}
		}

		// Enqueue next states
		for _, dir := range nextDirs {
			nextPos := curr.pos.Add(dir)
			if d.grid.inBounds(nextPos) {
				nextState := state{pos: nextPos, dir: dir}
				queue = append(queue, nextState)
			}
		}
	}

	return len(energized)
}

func (d *day16) Part1() int {
	return d.simulateBeams(
		state{pos: utils.Coord{X: 0, Y: 0}, dir: utils.Right},
		make(map[state]struct{}),
		make(map[utils.Coord]struct{}),
	)
}

func (d *day16) Part2() int {
	var (
		maxEnergized int
		energized    int
		visited      = make(map[state]struct{})
		energizedMap = make(map[utils.Coord]struct{})
	)

	for x := 0; x < len(d.grid[0]); x++ {
		energized = d.simulateBeams(
			state{pos: utils.Coord{X: x, Y: 0}, dir: utils.Down},
			visited,
			energizedMap,
		)
		maxEnergized = max(maxEnergized, energized)
		energized = d.simulateBeams(
			state{pos: utils.Coord{X: x, Y: len(d.grid) - 1}, dir: utils.Up},
			visited,
			energizedMap,
		)
		maxEnergized = max(maxEnergized, energized)
	}

	for y := 0; y < len(d.grid); y++ {
		energized = d.simulateBeams(
			state{pos: utils.Coord{X: 0, Y: y}, dir: utils.Right},
			visited,
			energizedMap,
		)
		maxEnergized = max(maxEnergized, energized)
		energized = d.simulateBeams(
			state{pos: utils.Coord{X: len(d.grid[0]) - 1, Y: y}, dir: utils.Left},
			visited,
			energizedMap,
		)
		maxEnergized = max(maxEnergized, energized)
	}

	return maxEnergized
}

func Parse(filename string) *day16 {
	data := utils.ReadLines(filename)
	grid := make(grid, len(data))

	for y, line := range data {
		grid[y] = make([]cell, len(line))
		for x, r := range line {
			val, ok := mapRune[r]
			if !ok {
				panic("invalid rune")
			}
			grid[y][x] = 1 << val
		}
	}

	return &day16{
		grid: grid,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of energized tiles:", day.Part1())
	fmt.Println("ANSWER2: max number of energized tiles possible:", day.Part2())
}
