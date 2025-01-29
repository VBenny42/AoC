package day16

import (
	"fmt"
	"sync"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	cell int
	grid [][]cell
)

type day16 struct {
	grid grid
}

const empty cell = 0
const (
	vSplit cell = 1 << iota
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

func (d *day16) simulateBeams(start state) int {
	var (
		queue     = []state{start}
		visited   = make(map[state]struct{})
		energized = make(map[utils.Coord]struct{})
	)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if _, ok := visited[curr]; ok {
			// Remove beam if it's already been visited
			// Loop will happen otherwise
			continue
		}
		visited[curr] = struct{}{}

		energized[curr.pos] = struct{}{}

		// Determine the next directions based on the current cell and direction
		cellType := d.grid[curr.pos.Y][curr.pos.X]

		var nextDirs []utils.Coord

		switch {
		case cellType == empty:
			nextDirs = []utils.Coord{curr.dir}
		case cellType&vSplit != 0:
			if curr.dir == utils.Up || curr.dir == utils.Down {
				// Treat like an empty cell
				nextDirs = []utils.Coord{curr.dir}
			} else {
				nextDirs = []utils.Coord{utils.Up, utils.Down}
			}
		case cellType&hSplit != 0:
			if curr.dir == utils.Left || curr.dir == utils.Right {
				// Treat like an empty cell
				nextDirs = []utils.Coord{curr.dir}
			} else {
				nextDirs = []utils.Coord{utils.Left, utils.Right}
			}
		case cellType&rightMirror != 0:
			switch curr.dir {
			case utils.Right:
				nextDirs = []utils.Coord{utils.Up}
			case utils.Left:
				nextDirs = []utils.Coord{utils.Down}
			case utils.Up:
				nextDirs = []utils.Coord{utils.Right}
			case utils.Down:
				nextDirs = []utils.Coord{utils.Left}
			}
		case cellType&leftMirror != 0:
			switch curr.dir {
			case utils.Right:
				nextDirs = []utils.Coord{utils.Down}
			case utils.Left:
				nextDirs = []utils.Coord{utils.Up}
			case utils.Up:
				nextDirs = []utils.Coord{utils.Left}
			case utils.Down:
				nextDirs = []utils.Coord{utils.Right}
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
	)
}

func (d *day16) Part2() int {
	var maxEnergized int

	var (
		wg          sync.WaitGroup
		energizedCh = make(chan int)
	)

	var (
		rowLen = len(d.grid[0]) - 1
		colLen = len(d.grid) - 1
	)
	for x := 0; x < len(d.grid[0]); x++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			energizedCh <- d.simulateBeams(
				state{pos: utils.Crd(x, 0), dir: utils.Down},
			)
		}()
		go func() {
			defer wg.Done()
			energizedCh <- d.simulateBeams(
				state{pos: utils.Crd(x, colLen), dir: utils.Down},
			)
		}()
	}

	for y := 0; y < len(d.grid); y++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			energizedCh <- d.simulateBeams(
				state{pos: utils.Crd(0, y), dir: utils.Right},
			)
		}()
		go func() {
			defer wg.Done()
			energizedCh <- d.simulateBeams(
				state{pos: utils.Crd(rowLen, y), dir: utils.Left},
			)
		}()
	}

	go func() {
		wg.Wait()
		close(energizedCh)
	}()

	for energized := range energizedCh {
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
			grid[y][x] = val
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
