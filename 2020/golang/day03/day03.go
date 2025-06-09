package day03

import (
	"fmt"
	"image"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

const (
	open = '.'
	tree = '#'
)

type day03 struct {
	grid utils.Grid[rune]
}

func (d *day03) countTrees(slope image.Point) (count int) {
	height := len(d.grid)

	for pos := image.Pt(0, 0); pos.Y < height; pos = pos.Add(slope) {
		if !d.grid.InBounds(pos) {
			pos.X = pos.X % len(d.grid[0])
		}
		if d.grid.Get(pos) == tree {
			count++
		}
	}

	return
}

func (d *day03) Part1And2() (count, product int) {
	slopes := []image.Point{
		image.Pt(1, 1),
		image.Pt(3, 1),
		image.Pt(5, 1),
		image.Pt(7, 1),
		image.Pt(1, 2),
	}

	product = 1
	for i, slope := range slopes {
		if i == 1 {
			count = d.countTrees(slope)
		}
		product *= d.countTrees(slope)
	}

	return count, product
}

func Parse(filename string) *day03 {
	lines := utils.ReadLines(filename)
	grid := utils.NewGrid[rune](len(lines[0]), len(lines))

	for y, line := range lines {
		for x, char := range line {
			if char != open && char != tree {
				panic(fmt.Sprintf("Invalid character '%c' in input", char))
			}
			grid.Set(image.Pt(x, y), char)
		}
	}

	return &day03{grid: grid}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println("ANSWER1: number of trees encountered:", part1)
	fmt.Println("ANSWER2: product of number of trees encountered on all slopes:", part2)
}
