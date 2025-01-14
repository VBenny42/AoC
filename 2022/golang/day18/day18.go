package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type cube struct {
	x, y, z int
}

type day18 struct {
	cubes            []cube
	grid             [][][]bool
	maxX, maxY, maxZ int
}

var directions = []cube{
	{x: 1, y: 0, z: 0},
	{x: -1, y: 0, z: 0},
	{x: 0, y: 1, z: 0},
	{x: 0, y: -1, z: 0},
	{x: 0, y: 0, z: 1},
	{x: 0, y: 0, z: -1},
}

func (d *day18) Part1() int {
	var coveredSurfaceArea int

	for _, cube := range d.cubes {
		for _, dir := range directions {
			neighbor := cube
			neighbor.x += dir.x
			neighbor.y += dir.y
			neighbor.z += dir.z
			if neighbor.x >= 0 && neighbor.x <= d.maxX &&
				neighbor.y >= 0 && neighbor.y <= d.maxY &&
				neighbor.z >= 0 && neighbor.z <= d.maxZ &&
				d.grid[neighbor.z][neighbor.y][neighbor.x] {
				coveredSurfaceArea++
			}
		}
	}

	return 6*len(d.cubes) - coveredSurfaceArea
}

func (d *day18) canReachEdge(c cube) bool {
	queue := []cube{c}
	seen := make(map[cube]bool)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if seen[curr] || (curr.x >= 0 && curr.x <= d.maxX &&
			curr.y >= 0 && curr.y <= d.maxY &&
			curr.z >= 0 && curr.z <= d.maxZ) &&
			d.grid[curr.z][curr.y][curr.x] {
			continue
		}
		seen[curr] = true

		// Check if we've reached an edge
		if curr.x <= 0 || curr.x >= d.maxX ||
			curr.y <= 0 || curr.y >= d.maxY ||
			curr.z <= 0 || curr.z >= d.maxZ {
			return true
		}

		// Add neighbors to queue
		for _, dir := range directions {
			next := cube{
				x: curr.x + dir.x,
				y: curr.y + dir.y,
				z: curr.z + dir.z,
			}
			queue = append(queue, next)
		}
	}
	return false
}

func (d *day18) Part2() int {
	externalSurfaceArea := 0

	for _, c := range d.cubes {
		for _, dir := range directions {
			next := cube{
				x: c.x + dir.x,
				y: c.y + dir.y,
				z: c.z + dir.z,
			}
			if d.canReachEdge(next) {
				externalSurfaceArea++
			}
		}
	}

	return externalSurfaceArea
}

func Parse(filename string) *day18 {
	data := utils.ReadLines(filename)
	cubes := make([]cube, len(data))
	maxX, maxY, maxZ := 0, 0, 0

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
	for col := range grid {
		grid[col] = make([][]bool, maxY+1)
		for row := range grid[col] {
			grid[col][row] = make([]bool, maxX+1)
		}
	}

	for _, c := range cubes {
		grid[c.z][c.y][c.x] = true
	}

	return &day18{
		cubes: cubes,
		grid:  grid,
		maxX:  maxX,
		maxY:  maxY,
		maxZ:  maxZ,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: total surface area:", day.Part1())
	fmt.Println("ANSWER2: exterior surface area:", day.Part2())
}
