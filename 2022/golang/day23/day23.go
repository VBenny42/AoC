package day23

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type day23 struct {
	bitmap bitmap
}

func (d *day23) Part1() int {
	return d.bitmap.rounds(10)
}

func (d *day23) Part2() int {
	return d.bitmap.rounds(10000000)
}

func Parse(filename string) *day23 {
	data := utils.ReadLines(filename)

	var bitmap bitmap

	for y, line := range data {
		for x, char := range line {
			if char == '#' {
				bitmap.set(utils.Coord{X: x, Y: y})
			}
		}
	}

	return &day23{bitmap}
}

func Solve(filename string) {
	fmt.Println("ANSWER1: number of empty ground tiles after 10 rounds:", Parse(filename).Part1())
	fmt.Println("ANSWER2: round where no elf moves:", Parse(filename).Part2())
}
