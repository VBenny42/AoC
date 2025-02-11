package day15

import (
	"fmt"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day15 struct {
	grid utils.Grid[int]
}

// Need to find path from top left to bottom right
// Can move up, down, left, right
// Need to find path with lowest sum of numbers
// Therefore, we can use Dijkstra's algorithm to find the shortest path
func (d *day15) Part1() int {
	return d.dijkstra()
}

// Extend grid 5 times to the right and 5 down
// Increase risk by 1 for each cell, wrapping around from 1 if it exceeds 9
// Call dijkstra again (and pray it doesn't take too long)
func (d *day15) Part2() int {
	newGrid := utils.NewGrid[int](len(d.grid[0])*5, len(d.grid)*5)
	height, width := len(d.grid), len(d.grid[0])

	for tileY := 0; tileY < 5; tileY++ {
		for tileX := 0; tileX < 5; tileX++ {
			for y, row := range d.grid {
				for x, risk := range row {
					// Calculate new risk level
					increase := tileX + tileY
					newRisk := risk + increase
					if newRisk > 9 {
						newRisk = newRisk - 9
					}
					// Calculate new position
					newX := x + (width * tileX)
					newY := y + (height * tileY)
					newGrid[newY][newX] = newRisk
				}
			}
		}
	}

	d.grid = newGrid

	return d.dijkstra()
}

func Parse(filename string) *day15 {
	var (
		data = utils.ReadLines(filename)
		grid = utils.NewGrid[int](len(data[0]), len(data))
	)

	for y, line := range data {
		for x, char := range line {
			grid[y][x] = int(char) - '0'
		}
	}

	return &day15{grid: grid}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println(
		"ANSWER1: lowest total risk of any path from top left to bottom right:",
		day.Part1(),
	)
	fmt.Println(
		"ANSWER2: lowest total risk of any path from top left to bottom right of extended grid:",
		day.Part2(),
	)
}
