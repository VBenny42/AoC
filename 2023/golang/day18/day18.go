package day18

import (
	"fmt"
	"image"
	"image/color"
	// "path/filepath"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type instruction struct {
	direction utils.Coord
	distance  int
	color     color.RGBA
	colorStr  string
}

type grid [][]color.RGBA

type day18 struct {
	instructions []instruction
	grid         grid
	origin       utils.Coord
}

// Borrowed from
// https://stackoverflow.com/questions/54197913/parse-hex-string-to-rgb-in-golang
func ParseHexColorFast(s string) (c color.RGBA) {
	c.A = 0xff

	if s[0] != '#' {
		return c
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		return 0
	}

	c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
	c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
	c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	return
}

// Assuming center spot of graph is always in interior
func (g *grid) countInterior() (sum int) {
	queue := []utils.Coord{{X: len((*g)[0]) / 2, Y: len(*g) / 2}}
	fill := color.RGBAModel.Convert(image.White).(color.RGBA)
	var sentinel color.RGBA

	// fill until we hit a colored edge
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if (*g)[curr.Y][curr.X] == fill {
			continue
		}
		(*g)[curr.Y][curr.X] = fill
		sum++

		for _, dir := range utils.Directions {
			next := curr.Add(dir)
			if next.X < 0 || next.X >= len((*g)[0]) || next.Y < 0 || next.Y >= len(*g) ||
				(*g)[next.Y][next.X] != sentinel {
				continue
			}
			queue = append(queue, next)
		}
	}

	return
}

func (d *day18) parseColorAsInstructions() {
	newInstructions := make([]instruction, len(d.instructions))

	for i, instr := range d.instructions {
		var newDir utils.Coord
		switch instr.colorStr[5] {
		case '0':
			newDir = utils.Right
		case '1':
			newDir = utils.Down
		case '2':
			newDir = utils.Left
		case '3':
			newDir = utils.Up
		}

		newDist, err := strconv.ParseUint(instr.colorStr[:5], 16, 64)
		if err != nil {
			panic(err)
		}

		newInstructions[i] = instruction{
			direction: newDir,
			distance:  int(newDist),
			color:     instr.color,
			colorStr:  instr.colorStr,
		}
	}

	d.instructions = newInstructions
}

func (d *day18) calculateTrenchWithShoelace() int {
	// Collect vertices of the polygon
	var (
		vertices  = make([]utils.Coord, len(d.instructions)+1)
		curr      = utils.Coord{X: 0, Y: 0}
		perimeter int64
	)
	vertices[0] = curr

	for i, instr := range d.instructions {
		// Calculate next vertex
		curr = utils.Coord{
			X: curr.X + instr.direction.X*instr.distance,
			Y: curr.Y + instr.direction.Y*instr.distance,
		}
		vertices[i+1] = curr
		perimeter += int64(instr.distance)
	}

	// Calculate area using Shoelace formula
	var area int64
	for i := 0; i < len(vertices)-1; i++ {
		area += int64(vertices[i].X)*int64(vertices[i+1].Y) -
			int64(vertices[i+1].X)*int64(vertices[i].Y)
	}
	area = utils.Abs(area) / 2

	// Use Pick's theorem to calculate total points
	// A = i + b/2 - 1 where:
	// A = area
	// i = interior points
	// b = boundary points (our perimeter)
	//
	// Rearranging to solve for total points (interior + boundary):
	// total = i + b = A + b/2 + 1
	return int(area + perimeter/2 + 1)
}

func (d *day18) Part1() (count int) {
	curr := d.origin
	for _, instr := range d.instructions {
		count += instr.distance
		for i := 0; i < instr.distance; i++ {
			d.grid[curr.Y][curr.X] = instr.color
			curr = curr.Add(instr.direction)
		}
	}

	count += d.grid.countInterior()

	return

	// Using shoelace is around 1.5ms faster on my laptop,
	// but I like the above method more :)
	// I get to make a bitmap of the grid and see the trench!
	// Visualizing the grid is a lot more fun than just calculating the area hehe

	// return d.calculateTrenchWithShoelace()
}

func (d *day18) Part2() int {
	d.parseColorAsInstructions()
	return d.calculateTrenchWithShoelace()
}

func Parse(filename string) *day18 {
	var (
		data         = utils.ReadLines(filename)
		instructions = make([]instruction, len(data))
	)

	var (
		minX, maxX   int
		minY, maxY   int
		currX, currY int
	)

	fieldsFunc := func(r rune) bool {
		return r == ' ' || r == '(' || r == ')'
	}

	for i, line := range data {
		var (
			fields    = strings.FieldsFunc(line, fieldsFunc)
			direction utils.Coord
			distance  = utils.Atoi(fields[1])
			color     = ParseHexColorFast(fields[2])
		)

		switch fields[0] {
		case "U":
			direction = utils.Up
		case "D":
			direction = utils.Down
		case "L":
			direction = utils.Left
		case "R":
			direction = utils.Right
		}

		instructions[i] = instruction{
			direction: direction,
			distance:  distance,
			color:     color,
			colorStr:  fields[2][1:],
		}

		dx := distance * direction.X
		dy := distance * direction.Y
		currX += dx
		currY += dy

		minX = min(minX, currX)
		minY = min(minY, currY)
		maxX = max(maxX, currX)
		maxY = max(maxY, currY)
	}

	width := maxX - minX + 1
	height := maxY - minY + 1

	grid := make(grid, height)
	for i := range grid {
		grid[i] = make([]color.RGBA, width)
	}

	return &day18{
		instructions: instructions,
		grid:         grid,
		origin:       utils.Coord{X: -minX, Y: -minY},
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: cubic metres of lava in trench:", day.Part1())

	// // Uncomment this to get bitmap of grid
	// bitmapFilename := filepath.Base(filename) + ".bmp"
	// if err := day.grid.writeBitmap(bitmapFilename); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Bitmap saved to", bitmapFilename)

	fmt.Println("ANSWER2: cubic metres of lava in trench after converting colors:", day.Part2())
}
