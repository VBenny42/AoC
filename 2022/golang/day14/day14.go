package day14

import (
	"fmt"
	"os"
	"regexp"
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

func (g *grid) printBitmap(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := *g
	line := strings.Builder{}
	line.WriteString(fmt.Sprintf("P2\n%d %d\n15\n", len(grid[0]), len(grid)))

	mapping := map[rune]int{
		'.': 0,
		'#': 15,
		'o': 7,
		'+': 10,
	}

	for _, row := range *g {
		for _, cell := range row {
			line.WriteString(fmt.Sprintf("%d ", mapping[cell]))
		}
		line.WriteString("\n")
	}

	_, err = file.WriteString(line.String())
	if err != nil {
		panic(err)
	}
}

func (d *Day14) fall(grain utils.Coord) error {
	// Check if starting position is blocked
	if d.Grid.at(grain) == '#' || d.Grid.at(grain) == 'o' {
		return fmt.Errorf("Obstacle")
	}

	// Try to move down
	next := grain.Add(utils.Down)
	if !d.inBounds(next) {
		return fmt.Errorf("Reached bottom")
	}
	if d.Grid.at(next) == '.' {
		err := d.fall(next)
		if err == nil {
			return nil // Sand was placed successfully below
		}
		if err.Error() == "Reached bottom" {
			return err
		}
		// Otherwise try left/right
	}

	// Hit obstacle, check left
	left := next.Add(utils.Left)
	if !d.inBounds(left) {
		return fmt.Errorf("Reached bottom")
	}
	if d.Grid.at(left) == '.' {
		err := d.fall(left)
		if err == nil {
			return nil // Sand was placed successfully to the left
		}
		if err.Error() == "Reached bottom" {
			return err
		}
	}

	// Check right
	right := next.Add(utils.Right)
	if !d.inBounds(right) {
		return fmt.Errorf("Reached bottom")
	}
	if d.Grid.at(right) == '.' {
		err := d.fall(right)
		if err == nil {
			return nil // Sand was placed successfully to the right
		}
		if err.Error() == "Reached bottom" {
			return err
		}
	}

	// Can't move down, left, or right - come to rest
	d.Grid[grain.Y][grain.X] = 'o'
	return nil
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
	pattern := `(\d+),(\d+)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(line, -1)

	lines := make([]Line, len(matches)-1)
	maxY := 0

	for i := 0; i < len(matches)-1; i++ {
		start, end := matches[i], matches[i+1]
		lines[i] = Line{
			Start: utils.Coord{X: utils.Must(strconv.Atoi(start[1])), Y: utils.Must(strconv.Atoi(start[2]))},
			End:   utils.Coord{X: utils.Must(strconv.Atoi(end[1])), Y: utils.Must(strconv.Atoi(end[2]))},
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

	// rocks will form a triangle, with point being sand source
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
	sourceX := 500 - minX
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
}
