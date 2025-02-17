package day09

import (
	"fmt"
	"image"
	"sort"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day09 struct {
	grid      utils.Grid[int]
	lowPoints []image.Point
}

func (d *day09) Part1() (sum int) {
	for y, row := range d.grid {
		for x, cell := range row {
			pt := image.Pt(x, y)

			higherThanNeighbors := false
			for _, neighbor := range d.grid.Neighbors(pt) {
				if cell >= d.grid.Get(neighbor) {
					higherThanNeighbors = true
					break
				}
			}
			if higherThanNeighbors {
				continue
			}

			sum += cell + 1
			d.lowPoints = append(d.lowPoints, pt)
		}
	}

	return
}

func (d *day09) Part2() int {
	var (
		visited = utils.NewGrid[bool](len(d.grid[0]), len(d.grid))
		sizes   = make([]int, 0, len(d.lowPoints))

		size  int
		curr  image.Point
		queue []image.Point
	)

	for _, lowPoint := range d.lowPoints {
		size = 0
		queue = queue[:0]
		queue = append(queue, lowPoint)

		// flood fill from lowPoint in every direction until 9 is reached

		for len(queue) > 0 {
			curr = queue[0]
			queue = queue[1:]

			if visited.Get(curr) {
				continue
			}
			visited.Set(curr, true)
			size++

			for _, neighbor := range d.grid.Neighbors(curr) {
				if d.grid.Get(neighbor) == 9 {
					continue
				}
				queue = append(queue, neighbor)
			}
		}

		sizes = append(sizes, size)
	}

	sort.Ints(sizes)

	length := len(sizes)

	return sizes[length-3] * sizes[length-2] * sizes[length-1]
}

func Parse(filename string) *day09 {
	var (
		data = utils.ReadLines(filename)
		grid = utils.NewGrid[int](len(data[0]), len(data))
	)

	for y, line := range data {
		for x, char := range line {
			grid[y][x] = int(char) - '0'
		}
	}

	return &day09{grid: grid}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of risk levels of low points on map:", day.Part1())
	fmt.Println("ANSWER2: product of the three largest basins:", day.Part2())

	// g := grid(day.grid)
	// g.writeImage(filename)
}
