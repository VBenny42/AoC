package day10

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type grid struct {
	grid [][]rune
	rows int
	cols int
}

type day10 struct {
	grid  grid
	start utils.Coord
}

var pipes = map[rune][2]utils.Coord{
	'|': {
		utils.Up, utils.Down,
	},
	'-': {
		utils.Left, utils.Right,
	},
	'L': {
		utils.Up, utils.Right,
	},
	'J': {
		utils.Up, utils.Left,
	},
	'7': {
		utils.Down, utils.Left,
	},
	'F': {
		utils.Down, utils.Right,
	},
}

func (d *day10) Part1And2() (startToFurthest int, enclosedTiles int) {
	var (
		visited = make(map[utils.Coord]struct{})
		queue   = []utils.Coord{d.start}
	)

	for len(queue) > 0 {
		curr := queue[0]

		// Found a loop
		if _, ok := visited[curr]; ok {
			break
		}
		visited[curr] = struct{}{}
		queue = queue[1:]

		directions := pipes[d.grid.grid[curr.Y][curr.X]]
		for _, dir := range directions {
			next := curr.Add(dir)
			_, ok := visited[next]
			if d.grid.inBounds(next) && !ok {
				queue = append(queue, next)
			}
		}
	}

	startToFurthest = len(visited) / 2

	// Part 2
	var expandedGrid grid

	// Make a new grid with space between each character
	expanded := make([][]rune, 0, d.grid.rows*2+2)
	expanded = append(expanded, make([]rune, d.grid.cols*2+2))
	for y, row := range d.grid.grid {
		expanded = append(expanded, make([]rune, d.grid.cols*2+2))
		for x, char := range row {
			if _, ok := visited[utils.Crd(x, y)]; ok {
				expanded[y*2+1][x*2+1] = char
			} else {
				expanded[y*2+1][x*2+1] = '.'
			}
		}
		expanded = append(expanded, make([]rune, d.grid.cols*2+2))
	}
	expandedGrid.grid = expanded
	expandedGrid.rows = len(expanded)
	expandedGrid.cols = len(expanded[0])

	// Try and connect pipes in empty spaces between pipes
	for y, row := range expanded {
		for x, char := range row {
			if char == 0 {
				expandedGrid.markPipe(utils.Crd(x, y))
			}
		}
	}

	queue = queue[:0]
	queue = append(queue, utils.Crd(0, 0))
	clear(visited)

	// Flood fill through grid,
	// marking all visited tiles as empty
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		_, ok := visited[curr]
		if ok {
			continue
		}
		visited[curr] = struct{}{}

		if expanded[curr.Y][curr.X] == '.' {
			expanded[curr.Y][curr.X] = 0
		}

		for _, dir := range utils.Directions {
			next := curr.Add(dir)
			if expandedGrid.inBounds(next) &&
				(expanded[next.Y][next.X] == '.' ||
					expanded[next.Y][next.X] == 0) {
				queue = append(queue, next)
			}
		}
	}

	// Count the number of tiles not touched by the flood fill
	for _, row := range expanded {
		for _, char := range row {
			if char == '.' {
				enclosedTiles++
			}
		}
	}

	return
}

func Parse(filename string) *day10 {
	var (
		data      = utils.ReadLines(filename)
		gridSlice = make([][]rune, len(data))
		start     utils.Coord
	)

	for y, line := range data {
		gridSlice[y] = []rune(line)
		for x, char := range line {
			if char == 'S' {
				start = utils.Crd(x, y)
			}
		}
	}

	grid := grid{grid: gridSlice, rows: len(gridSlice), cols: len(gridSlice[0])}
	grid.markPipe(start)

	return &day10{
		grid:  grid,
		start: start,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	startToFurthest, tilesEnclosed := day.Part1And2()

	fmt.Println("ANSWER1: distance between start and furthest point in loop:",
		startToFurthest)
	fmt.Println("ANSWER2: number of tiles enclosed by the loop:", tilesEnclosed)
}
