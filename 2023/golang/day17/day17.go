package day17

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type grid [][]int

type day17 struct {
	grid grid
}

func (g *grid) inBounds(c utils.Coord) bool {
	return c.X >= 0 && c.X < len((*g)[0]) && c.Y >= 0 && c.Y < len(*g)
}

type state struct {
	pos      utils.Coord
	dir      int // 0: north, 1: east, 2: south, 3: west
	steps    int
	heatLoss int
}

func (d *day17) Part1() int {
	return d.dijkstra1()
}

func (d *day17) Part2() int {
	return d.dijkstra2()
}

func Parse(filename string) *day17 {
	var (
		data = utils.ReadLines(filename)
		g    = make(grid, len(data))
	)

	for y, line := range data {
		g[y] = make([]int, len(line))
		for x, char := range line {
			g[y][x] = utils.Atoi(string(char))
		}
	}

	return &day17{grid: g}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: least heat loss incurred:", day.Part1())
	fmt.Println("ANSWER2: least heat loss incurred for ultra crucible:", day.Part2())
}
