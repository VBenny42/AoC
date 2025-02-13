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

type droplet byte

const (
	air droplet = iota
	lava
	steam
)

type day18 struct {
	cubes            []cube
	grid             *[][][]droplet
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

func (d *day18) floodFillWithSteam() {
	var start cube

	queue := []cube{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if (*d.grid)[curr.z][curr.y][curr.x] != air {
			continue
		}
		(*d.grid)[curr.z][curr.y][curr.x] = steam

		for _, dir := range directions {
			neighbor := cube{
				x: curr.x + dir.x,
				y: curr.y + dir.y,
				z: curr.z + dir.z,
			}

			isValid := neighbor.x >= 0 && neighbor.x <= d.maxX+1 &&
				neighbor.y >= 0 && neighbor.y <= d.maxY+1 &&
				neighbor.z >= 0 && neighbor.z <= d.maxZ+1

			if isValid && (*d.grid)[neighbor.z][neighbor.y][neighbor.x] == air {
				queue = append(queue, neighbor)
			}
		}
	}
}

func (d *day18) Part1() int {
	var coveredArea int

	for _, c := range d.cubes {
		for _, dir := range directions {
			neighbor := cube{
				x: c.x + dir.x,
				y: c.y + dir.y,
				z: c.z + dir.z,
			}

			isValid := neighbor.x >= 0 && neighbor.x <= d.maxX &&
				neighbor.y >= 0 && neighbor.y <= d.maxY &&
				neighbor.z >= 0 && neighbor.z <= d.maxZ

			if isValid && (*d.grid)[neighbor.z][neighbor.y][neighbor.x] == lava {
				coveredArea++
			}
		}
	}

	return 6*len(d.cubes) - coveredArea
}

func (d *day18) Part2() int {
	d.floodFillWithSteam()

	var exteriorArea int

	for _, c := range d.cubes {
		for _, dir := range directions {
			neighbor := cube{
				x: c.x + dir.x,
				y: c.y + dir.y,
				z: c.z + dir.z,
			}

			isValid := neighbor.x >= 0 && neighbor.x <= d.maxX+1 &&
				neighbor.y >= 0 && neighbor.y <= d.maxY+1 &&
				neighbor.z >= 0 && neighbor.z <= d.maxZ+1

			if isValid && (*d.grid)[neighbor.z][neighbor.y][neighbor.x] == steam {
				exteriorArea++
			}
		}
	}

	return exteriorArea
}

func Parse(filename string) *day18 {
	data := utils.ReadLines(filename)
	cubes := make([]cube, len(data))
	maxX, maxY, maxZ := 0, 0, 0

	for i, line := range data {
		split := strings.Split(line, ",")
		// Adding padding for steam to reach cubes on the edges
		cubes[i] = cube{
			x: utils.Must(strconv.Atoi(split[0])) + 1,
			y: utils.Must(strconv.Atoi(split[1])) + 1,
			z: utils.Must(strconv.Atoi(split[2])) + 1,
		}
		maxX = max(maxX, cubes[i].x)
		maxY = max(maxY, cubes[i].y)
		maxZ = max(maxZ, cubes[i].z)
	}

	// Adding padding for steam to fill all around the cubes
	grid := make([][][]droplet, maxZ+2)
	for col := range grid {
		grid[col] = make([][]droplet, maxY+2)
		for row := range grid[col] {
			grid[col][row] = make([]droplet, maxX+2)
		}
	}

	for _, c := range cubes {
		grid[c.z][c.y][c.x] = lava
	}

	return &day18{
		cubes: cubes,
		grid:  &grid,
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
