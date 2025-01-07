package day08

import (
	"fmt"
	"slices"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type grid [][]int

type day08 struct {
	grid grid
}

func (g *grid) getAdjacentCells(c utils.Coord) [][]int {
	grid := *g

	west := grid[c.X][:c.Y]
	east := grid[c.X][c.Y+1:]
	north := make([]int, c.X)
	south := make([]int, len(grid)-c.X-1)

	for i := 0; i < c.X; i++ {
		north[i] = grid[i][c.Y]
	}
	for i := c.X + 1; i < len(grid); i++ {
		south[i-c.X-1] = grid[i][c.Y]
	}

	return [][]int{west, east, north, south}
}

func (g *grid) isVisible(c utils.Coord) bool {
	treeHeight := (*g)[c.X][c.Y]

	adjacentCells := g.getAdjacentCells(c)

	for _, side := range adjacentCells {
		if len(side) == 0 || treeHeight > slices.Max(side) {
			return true
		}
	}

	return false
}

func (d *grid) getScore(c utils.Coord) int {
	treeHeight := (*d)[c.X][c.Y]

	adjacentCells := d.getAdjacentCells(c)
	slices.Reverse(adjacentCells[0])
	slices.Reverse(adjacentCells[2])

	getVisibleTrees := func(side []int) int {
		score := 0
		for _, height := range side {
			score++
			if height >= treeHeight {
				break
			}
		}
		return score
	}

	score := 1

	for _, side := range adjacentCells {
		score *= getVisibleTrees(side)
	}

	return score
}

func (d *day08) Part1and2() (int, int) {
	total, maxScore := 0, 0
	for y, row := range d.grid {
		for x := range row {
			if d.grid.isVisible(utils.Coord{X: x, Y: y}) {
				total++
			}
			score := d.grid.getScore(utils.Coord{X: x, Y: y})
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return total, maxScore
}

func Parse(filename string) *day08 {
	data := utils.ReadLines(filename)

	grid := make(grid, len(data))
	for i, line := range data {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}

	return &day08{grid}
}

func Solve(filename string) {
	day := Parse(filename)
	total, maxScore := day.Part1and2()
	fmt.Println("ANSWER1: trees visible outside of the grid:", total)
	fmt.Println("ANSWER2: best spot to build a treehouse:", maxScore)
}
