package day10

import (
	"fmt"
	"slices"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day10 struct {
	chargers []int
}

func (d *day10) Part1() int {
	var (
		diffs = make(map[int]int, 2)
		prev  = 0
	)

	for _, charger := range d.chargers {
		diff := charger - prev
		diffs[diff]++
		prev = charger
	}

	// Account for the device's built-in adapter
	diffs[3]++

	return diffs[1] * diffs[3]
}

func (d *day10) Part2() int {
	allJolts := append([]int{0}, d.chargers...)

	paths := make([]int, len(allJolts))

	// The first adapter (0 jolts) has only one way to be reached
	paths[0] = 1

	for i := 1; i < len(allJolts); i++ {
		for j := 0; j < i; j++ {
			if allJolts[i]-allJolts[j] <= 3 {
				paths[i] += paths[j]
			}
		}
	}

	return paths[len(paths)-1]
}

func Parse(filename string) *day10 {
	var (
		lines    = utils.ReadLines(filename)
		chargers = make([]int, len(lines))
		charger  int
	)

	for i, line := range lines {
		fmt.Sscanf(line, "%d", &charger)
		chargers[i] = charger
	}

	slices.Sort(chargers)

	return &day10{chargers: chargers}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: product of 1-jolt and 3-jolt differences:", day.Part1())
	fmt.Println("ANSWER2: number of distinct arrangements:", day.Part2())
}
