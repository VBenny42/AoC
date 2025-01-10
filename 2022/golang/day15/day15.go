package day15

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type pair struct {
	sensor utils.Coord
	beacon utils.Coord
}

type grid [][]rune

type day15 struct {
	pairs      []pair
	rowToCheck int
	grid       grid
}

func (g grid) String() string {
	str := "\n"
	for _, row := range g {
		str += string(row) + "\n"
	}
	return str
}

func manhattanDistance(a, b utils.Coord) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func (g *grid) manhattanNeighbours(p utils.Coord, dist int) []utils.Coord {
	neighbours := make([]utils.Coord, 0)
	width := len((*g)[0])
	height := len(*g)

	for dx := -dist; dx <= dist; dx++ {
		maxDY := dist - utils.Abs(dx) // This is correct from your original code
		for dy := -maxDY; dy <= maxDY; dy++ {
			newX := p.X + dx
			newY := p.Y + dy

			// Check bounds
			if newX >= 0 && newX < width && newY >= 0 && newY < height {
				neighbours = append(neighbours, utils.Coord{X: newX, Y: newY})
			}
		}
	}
	return neighbours
}

func (g *grid) fill(p pair) {
	dist := manhattanDistance(p.sensor, p.beacon)

	for _, position := range g.manhattanNeighbours(p.sensor, dist) {
		if (*g)[position.Y][position.X] == '.' {
			(*g)[position.Y][position.X] = 'X'
		}
	}
}

func (d *day15) Part1() int {
	for _, pair := range d.pairs {
		d.grid.fill(pair)
	}

	count := 0
	for _, cell := range d.grid[d.rowToCheck] {
		if cell == 'X' {
			count++
		}
	}

	return count
}

func Parse(filename string, rowToCheck int) *day15 {
	data := utils.ReadLines(filename)

	pairs := make([]pair, len(data))

	var maxRow, maxCol int
	minRow, minCol := 1000, 1000

	for i, line := range data {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&pairs[i].sensor.X, &pairs[i].sensor.Y, &pairs[i].beacon.X, &pairs[i].beacon.Y)
		maxRow = max(maxRow, pairs[i].sensor.Y, pairs[i].beacon.Y)
		maxCol = max(maxCol, pairs[i].sensor.X, pairs[i].beacon.X)
		minRow = min(minRow, pairs[i].sensor.Y, pairs[i].beacon.Y)
		minCol = min(minCol, pairs[i].sensor.X, pairs[i].beacon.X)
	}

	width := maxCol - minCol + 1
	height := maxRow - minRow + 1

	grid := make([][]rune, height)
	for y := range grid {
		grid[y] = make([]rune, width)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	for i := range pairs {
		pairs[i].sensor.X -= minCol
		pairs[i].sensor.Y -= minRow
		pairs[i].beacon.X -= minCol
		pairs[i].beacon.Y -= minRow
		grid[pairs[i].sensor.Y][pairs[i].sensor.X] = 'S'
		grid[pairs[i].beacon.Y][pairs[i].beacon.X] = 'B'
	}

	return &day15{pairs, rowToCheck, grid}
}

func Solve(filename string) {
	day := Parse(filename, 2000000)

	fmt.Println("ANSWER1: number of positions with no beacon in row 2000000:", day.Part1())
}
