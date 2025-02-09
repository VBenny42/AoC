package day11

import (
	"fmt"
	"image"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day11 struct {
	grid utils.Grid[int]
}

func (d *day11) step() (flashes int) {
	var queue []image.Point

	for y, row := range d.grid {
		for x := range row {
			d.grid[y][x]++
			if d.grid[y][x] > 9 {
				queue = append(queue, image.Pt(x, y))
			}
		}
	}

	seen := utils.NewGrid[bool](len(d.grid[0]), len(d.grid))
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if seen[curr.Y][curr.X] {
			continue
		}
		seen[curr.Y][curr.X] = true

		for _, dir := range utils.AllDirections {
			next := curr.Add(dir)
			if !d.grid.InBounds(next) {
				continue
			}

			d.grid[next.Y][next.X]++
			if d.grid[next.Y][next.X] > 9 {
				queue = append(queue, next)
			}
		}
	}

	for y, row := range seen {
		for x := range row {
			if seen[y][x] {
				d.grid[y][x] = 0
				flashes++
			}
		}
	}

	return
}

func (d *day11) Part1And2() (flashes, index int) {
	steps := 100
	for range steps {
		flashes += d.step()
	}

	goal := len(d.grid) * len(d.grid[0])
	var part2Flashes int

	for index = steps; part2Flashes != goal; index++ {
		part2Flashes = d.step()
	}

	return
}

func Parse(filename string) *day11 {
	var (
		data = utils.ReadLines(filename)
		grid = utils.NewGrid[int](len(data[0]), len(data))
	)

	for y, line := range data {
		for x, char := range line {
			grid[y][x] = int(char) - '0'
		}
	}

	return &day11{grid: grid}
}

func Solve(filename string) {
	day := Parse(filename)

	flashes, index := day.Part1And2()

	fmt.Println("ANSWER1: total flashes after 100 steps:", flashes)
	fmt.Println("ANSWER2: index where whole grid flashes:", index)
}
