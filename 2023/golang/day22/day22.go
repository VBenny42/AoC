package day22

import (
	"fmt"
	"slices"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	point struct {
		x, y, z int
	}
	brick struct {
		start, end point
	}
)

type day22 struct {
	bricks []brick
}

func newZCoords(zSlice [][]int, a, b point) (int, int) {
	var z1, z2 int

	switch {
	case a.x == b.x && a.z == b.z:
		for i := a.y; i <= b.y; i++ {
			z1 = max(z1, zSlice[i][a.x])
		}
		z1, z2 = z1+1, z1+1
		for i := a.y; i <= b.y; i++ {
			zSlice[i][a.x] = z1
		}

	case a.y == b.y && a.z == b.z:
		for i := a.x; i <= b.x; i++ {
			z1 = max(z1, zSlice[a.y][i])
		}
		z1, z2 = z1+1, z1+1
		for i := a.x; i <= b.x; i++ {
			zSlice[a.y][i] = z1
		}

	case a.x == b.x && a.y == b.y:
		z1 = zSlice[a.y][a.x] + 1
		z2 = z1 + (b.z - a.z)
		zSlice[a.y][a.x] = z2
	}

	return z1, z2
}

func blocksShifted(bricks []brick) (count int) {
	zSlice := make([][]int, 10)
	for i := range zSlice {
		zSlice[i] = make([]int, 10)
	}

	for i, b := range bricks {
		b.start.z, b.end.z = newZCoords(zSlice, b.start, b.end)

		if b.start.z != bricks[i].start.z || b.end.z != bricks[i].end.z {
			bricks[i].start.z = b.start.z
			bricks[i].end.z = b.end.z
			count++
		}
	}

	return
}

func (d *day22) Part1And2() (part1, part2 int) {
	blocksShifted(d.bricks)

	for i := range d.bricks {
		newBricks := make([]brick, len(d.bricks))
		copy(newBricks, d.bricks)

		newBricks = append(newBricks[:i], newBricks[i+1:]...)

		changes := blocksShifted(newBricks)
		if changes == 0 {
			part1++
		}
		part2 += changes
	}

	return
}

func Parse(filename string) *day22 {
	var (
		data   = utils.ReadLines(filename)
		bricks = make([]brick, len(data))
	)

	for i, line := range data {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == '~'
		})

		bricks[i].start = point{
			utils.Atoi(fields[0]),
			utils.Atoi(fields[1]),
			utils.Atoi(fields[2]),
		}
		bricks[i].end = point{
			utils.Atoi(fields[3]),
			utils.Atoi(fields[4]),
			utils.Atoi(fields[5]),
		}
	}

	slices.SortFunc(bricks, func(a, b brick) int {
		return a.start.z - b.start.z
	})

	return &day22{bricks}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println(
		"ANSWER1: number of bricks that can be disintegrated without causing any other bricks to fall:",
		part1,
	)
	fmt.Println(
		"ANSWER2: number of bricks that will fall due to chain reaction:",
		part2,
	)
}
