package day05

import (
	"fmt"
	"image"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type line struct {
	start, end image.Point
}

type grid utils.Grid[int]

type day05 struct {
	lines []line
	maxX  int
	maxY  int
	grid  grid
}

func (d *day05) Part1() (sum int) {
	d.grid = grid(utils.NewGrid[int](d.maxX+1, d.maxY+1))

	for _, l := range d.lines {
		if l.start.X != l.end.X && l.start.Y != l.end.Y {
			continue
		}
		var (
			startY, endY = l.start.Y, l.end.Y
			startX, endX = l.start.X, l.end.X
		)
		if startY > endY {
			startY, endY = endY, startY
		}
		if startX > endX {
			startX, endX = endX, startX
		}

		for y := startY; y <= endY; y++ {
			for x := startX; x <= endX; x++ {
				d.grid[y][x]++
			}
		}
	}

	for y := 0; y <= d.maxY; y++ {
		for x := 0; x <= d.maxX; x++ {
			if d.grid[y][x] > 1 {
				sum++
			}
		}
	}

	return
}

func (d *day05) Part2() (sum int) {
	// Now need to check diagonals
	for _, l := range d.lines {
		if l.start.X == l.end.X || l.start.Y == l.end.Y {
			continue
		}

		x, y := l.start.X, l.start.Y
		for {
			d.grid[y][x]++
			if x == l.end.X && y == l.end.Y {
				break
			}
			if x < l.end.X {
				x++
			} else {
				x--
			}
			if y < l.end.Y {
				y++
			} else {
				y--
			}
		}
	}

	for y := 0; y <= d.maxY; y++ {
		for x := 0; x <= d.maxX; x++ {
			if d.grid[y][x] > 1 {
				sum++
			}
		}
	}

	return
}

func Parse(filename string) *day05 {
	var (
		data  = utils.ReadLines(filename)
		lines = make([]line, len(data))
		maxX  int
		maxY  int
	)

	for i, line := range data {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == '-' || r == '>'
		})
		if len(fields) != 2 {
			panic(fmt.Sprintf("invalid line: %s", line))
		}
		left := strings.SplitN(fields[0], ",", 2)
		right := strings.SplitN(fields[1], ",", 2)

		lines[i].start = image.Pt(utils.Atoi(left[0]), utils.Atoi(left[1]))
		lines[i].end = image.Pt(utils.Atoi(right[0]), utils.Atoi(right[1]))
		maxX = max(maxX, lines[i].start.X, lines[i].end.X)
		maxY = max(maxY, lines[i].start.Y, lines[i].end.Y)
	}

	return &day05{
		lines: lines,
		maxX:  maxX,
		maxY:  maxY,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of points where at least two lines overlap:", day.Part1())
	fmt.Println(
		"ANSWER2: number of points where at least two lines overlap with diagonals:",
		day.Part2(),
	)

	// day.grid.writeImage("day05.png")
}
