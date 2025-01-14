package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type (
	cube struct {
		x, y, z int
	}
	grid [][][]bool
)

type day18 struct {
	grid  grid
	cubes []cube
}

// directions represents the 6 adjacent faces of a cube
var directions = []cube{
	{x: 1, y: 0, z: 0},
	{x: -1, y: 0, z: 0},
	{x: 0, y: 1, z: 0},
	{x: 0, y: -1, z: 0},
	{x: 0, y: 0, z: 1},
	{x: 0, y: 0, z: -1},
}

func (g *grid) neighbors(c cube) []cube {
	var neighbors []cube
	width, height, depth := len((*g)[0][0]), len((*g)[0]), len(*g)

	for _, dir := range directions {
		newX, newY, newZ := c.x+dir.x, c.y+dir.y, c.z+dir.z

		if newX < 0 || newY < 0 || newZ < 0 ||
			newX >= width || newY >= height || newZ >= depth {
			continue
		}

		if (*g)[newZ][newY][newX] {
			neighbors = append(neighbors, cube{newX, newY, newZ})
		}
	}

	return neighbors
}

func (d *day18) Part1() int {
	var surfaceArea int

	for _, cube := range d.cubes {
		surfaceArea += 6 - len(d.grid.neighbors(cube))
	}

	return surfaceArea
}

func Parse(filename string) *day18 {
	data := utils.ReadLines(filename)

	maxX, maxY, maxZ := 0, 0, 0

	cubes := make([]cube, len(data))
	for i, line := range data {
		split := strings.Split(line, ",")
		cubes[i] = cube{
			x: utils.Must(strconv.Atoi(split[0])),
			y: utils.Must(strconv.Atoi(split[1])),
			z: utils.Must(strconv.Atoi(split[2])),
		}
		maxX = max(maxX, cubes[i].x)
		maxY = max(maxY, cubes[i].y)
		maxZ = max(maxZ, cubes[i].z)
	}

	grid := make([][][]bool, maxZ+1)
	for z := range grid {
		grid[z] = make([][]bool, maxY+1)
		for y := range grid[z] {
			grid[z][y] = make([]bool, maxX+1)
		}
	}

	for _, cube := range cubes {
		grid[cube.z][cube.y][cube.x] = true
	}

	return &day18{cubes: cubes, grid: grid}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: total surface area:", day.Part1())
}
