package day14

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type Line struct {
	Start utils.Coord
	End   utils.Coord
}

type grid [][]rune

type Day14 struct {
	Lines  []Line
	Grid   grid
	source utils.Coord
	width  int
	height int
}

func (g grid) String() string {
	str := "\n"
	for _, row := range g {
		str += string(row) + "\n"
	}
	return str
}

func (d *Day14) inBounds(c utils.Coord) bool {
	return c.X >= 0 && c.X < d.width && c.Y >= 0 && c.Y < d.height
}

func (g *grid) at(c utils.Coord) rune {
	return (*g)[c.Y][c.X]
}

func (g *grid) makeBitmap(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := *g
	var line strings.Builder
	line.WriteString(fmt.Sprintf("P2\n%d %d\n15\n", len(grid[0]), len(grid)))

	mapping := map[rune]string{
		'.': "0",
		'#': "15",
		'o': "7",
		'+': "10",
	}

	for _, row := range *g {
		for _, cell := range row {
			line.WriteString(mapping[cell])
			line.WriteString(" ")
		}
		line.WriteString("\n")
	}

	_, err = file.WriteString(line.String())
	if err != nil {
		panic(err)
	}
}

func (d *Day14) fall(grain utils.Coord) error {
	for {
		if d.Grid.at(grain) == '#' || d.Grid.at(grain) == 'o' {
			return fmt.Errorf("Obstacle")
		}

		next := grain.Add(utils.Down)
		if !d.inBounds(next) {
			return fmt.Errorf("Reached bottom")
		}

		if d.Grid.at(next) != '.' {
			left := next.Add(utils.Left)
			right := next.Add(utils.Right)

			// check if we'll come to rest immediately
			if !d.inBounds(left) || !d.inBounds(right) ||
				(d.Grid.at(left) != '.' && d.Grid.at(right) != '.') {
				d.Grid[grain.Y][grain.X] = 'o'
				return nil
			}

			// try left
			if d.inBounds(left) && d.Grid.at(left) == '.' {
				grain = left
				continue
			}

			// try right
			if d.inBounds(right) && d.Grid.at(right) == '.' {
				grain = right
				continue
			}
		}

		// move down
		grain = next
	}
}

func (d *Day14) fillGrid() int {
	grains := 0
	var err error
	for {
		err = d.fall(d.source)
		if err != nil {
			break
		}
		grains++
	}
	return grains
}

func (d *Day14) Part1And2() (int, int) {
	part1 := d.fillGrid()
	// change height and width to be the entire Grid
	d.height = len(d.Grid)
	d.width = len(d.Grid[0])
	// Fill grid as much as possible
	part2 := part1 + d.fillGrid()

	return part1, part2
}

func ParseLine(line string) ([]Line, int) {
	matches := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == '-' || r == '>'
	})

	matchesSplit := make([][]string, len(matches))
	for i := range matches {
		matchesSplit[i] = strings.Split(matches[i], ",")
	}

	lines := make([]Line, len(matches)-1)
	maxY := 0
	for i := range lines {
		start, end := matchesSplit[i], matchesSplit[i+1]
		lines[i] = Line{
			Start: utils.Coord{
				X: utils.Must(strconv.Atoi(start[0])),
				Y: utils.Must(strconv.Atoi(start[1])),
			},
			End: utils.Coord{
				X: utils.Must(strconv.Atoi(end[0])),
				Y: utils.Must(strconv.Atoi(end[1])),
			},
		}
		maxY = max(maxY, lines[i].Start.Y, lines[i].End.Y)
	}

	return lines, maxY
}

func Parse(filename string) *Day14 {
	data := utils.ReadLines(filename)
	lines := make([]Line, 0)

	// First pass: find boundaries
	maxY := 0 // Initialize with sand source point
	for _, line := range data {
		parsedLine, lineY := ParseLine(line)
		lines = append(lines, parsedLine...)
		maxY = max(maxY, lineY)
	}

	part1Width := 2*maxY + 1
	part1Height := maxY + 1

	maxY += 2

	// Create grid with correct dimensions
	width := 2*maxY + 1
	grid := make(grid, maxY+1)
	for i := range grid {
		grid[i] = make([]rune, width)
		// Initialize with air
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// rocks will form a triangle, with top point being sand source
	// the width of the triangle will be 2*maxY + 1
	// therefore, the leftmost rock will be at 500 - maxY
	// no sand can be placed to the left of this point

	minX := 500 - maxY

	// Draw rocks
	for _, line := range lines {
		startX := line.Start.X - minX
		endX := line.End.X - minX
		startY := line.Start.Y
		endY := line.End.Y

		// Ensure we draw from smaller to larger coordinates
		if startX > endX {
			startX, endX = endX, startX
		}
		if startY > endY {
			startY, endY = endY, startY
		}

		// Draw the line
		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				grid[y][x] = '#'
			}
		}
	}

	for i := range grid[maxY] {
		grid[maxY][i] = '#'
	}

	// Mark the sand source
	sourceX := maxY
	grid[0][sourceX] = '+'

	return &Day14{
		Lines:  lines,
		Grid:   grid,
		source: utils.Coord{X: sourceX, Y: 0},
		height: part1Height,
		width:  part1Width,
	}
}

func Solve(filename string) {
	d := Parse(filename)

	part1, part2 := d.Part1And2()

	fmt.Println("ANSWER1: number of grains before infinite fall:", part1)
	fmt.Println("ANSWER2: number of grains that will fall in total:", part2)
	// d.Grid.printBitmap("day14.pgm")
}
