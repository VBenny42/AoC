package day14

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	cell rune
	grid [][]cell
)

type day14 struct {
	grid grid
}

const (
	empty   cell = '.'
	rounded cell = 'O'
	cube    cell = '#'
)

func (d *day14) Part1() (sum int) {
	d.grid.moveAllRocksUp()
	return d.grid.calculateNorthLoad()
}

func (d *day14) Part2() (sum int) {
	var (
		cycleFound bool
		i          int
		limit      = 1000000000
		seen       = make(map[uint32]int)
	)

	for i < limit {
		// go through a cycle
		for range 4 {
			d.grid.moveAllRocksUp()
			d.grid.rotateClockWise()
		}

		if !cycleFound {
			hash := d.grid.hash()
			if cycleStart, ok := seen[hash]; ok {
				cycleLen := i - cycleStart
				repeat := (limit - i) / cycleLen
				// Advance i as much as possible until limit
				i += repeat * cycleLen
				cycleFound = true
			} else {
				seen[hash] = i
			}
		}

		i++
	}

	return d.grid.calculateNorthLoad()
}

func Parse(filename string) *day14 {
	var (
		data    = utils.ReadLines(filename)
		gridVar = make(grid, len(data))
	)

	for y, line := range data {
		gridVar[y] = make([]cell, len(line))
		for x, c := range line {
			gridVar[y][x] = cell(c)
			if gridVar[y][x] != rounded && gridVar[y][x] != cube && gridVar[y][x] != empty {
				panic(fmt.Sprintf("unexpected cell: %c", gridVar[y][x]))
			}
		}
	}

	return &day14{
		grid: gridVar,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: total lood on north support beams:", day.Part1())
	fmt.Println("ANSWER2: total lood on north support beams after 1_000_000_000 cycles:",
		day.Part2())
}
