package day20

import (
	"fmt"
	"image"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type (
	pixel int
	grid  = utils.Grid[pixel]
)

const (
	dark pixel = iota
	light
)

type day20 struct {
	algorithm     []pixel
	grid          grid
	infinitePixel pixel
}

func getSquareIndex(g grid, p image.Point, infinitePixel pixel) (index int) {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			p := image.Point{p.X + x, p.Y + y}
			if g.InBounds(p) {
				if g[p.Y][p.X] == light {
					index += 1 << ((2-(y+1))*3 + (2 - (x + 1)))
				}
			} else {
				if infinitePixel == light {
					index += 1 << ((2-(y+1))*3 + (2 - (x + 1)))
				}
			}
		}
	}

	return
}

func (d *day20) enhance() (count int) {
	enhanced := utils.NewGrid[pixel](len(d.grid[0])+2, len(d.grid)+2)

	for y, row := range enhanced {
		for x := range row {
			index := getSquareIndex(d.grid, image.Point{x - 1, y - 1}, d.infinitePixel)
			enhanced[y][x] = d.algorithm[index]
			count += int(enhanced[y][x])
		}
	}

	if d.infinitePixel == dark {
		d.infinitePixel = d.algorithm[0]
	} else {
		d.infinitePixel = d.algorithm[511]
	}

	d.grid = enhanced

	return
}

func (d *day20) Part1() int {
	d.enhance()
	return d.enhance()
}

func (d *day20) Part2() int {
	for i := 2; i < 49; i++ {
		d.enhance()
	}
	return d.enhance()
}

// Things to note:
// - Every time the image is enhanced,
// size of grid should increase by 2 in width and height
// - 5x5 -> 7x7 -> 9x9 -> 11x11 -> ...
// Image is always square?
// Should I use a 2D array, or a map?
// Use array first, depending on part 2, I might change it to a map
// Array was better

func Parse(filename string) *day20 {
	var (
		data      = utils.ReadLines(filename)
		algorithm = make([]pixel, len(data[0]))
		grid      = utils.NewGrid[pixel](len(data[2]), len(data[2:]))
	)
	for i, char := range data[0] {
		switch char {
		case '#':
			algorithm[i] = light
		case '.':
			algorithm[i] = dark
		default:
			panic("Invalid character in algorithm")
		}
	}

	for y, line := range data[2:] {
		for x, char := range line {
			switch char {
			case '#':
				grid[y][x] = light
			case '.':
				grid[y][x] = dark
			default:
				panic("Invalid character in grid")
			}
		}
	}

	return &day20{algorithm: algorithm, grid: grid, infinitePixel: dark}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of tiles lit in the resulting image:", day.Part1())
	fmt.Println("ANSWER2: number of tiles lit after 50 enhancements:", day.Part2())

	// if err := imageGrid(day.grid).writeImage("day20.png"); err != nil {
	// 	fmt.Println("Error writing image:", err)
	// }
}
