package day01

import (
	"fmt"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day01 struct {
	depths []int
}

func (d *day01) Part1() (sum int) {
	current := d.depths[0]
	for _, depth := range d.depths[1:] {
		if depth > current {
			sum++
		}
		current = depth
	}

	return
}

func (d *day01) Part2() (sum int) {
	for i := 0; i < len(d.depths)-3; i++ {
		currSum := d.depths[i] + d.depths[i+1] + d.depths[i+2]
		nextSum := d.depths[i+1] + d.depths[i+2] + d.depths[i+3]
		if nextSum > currSum {
			sum++
		}
	}

	return
}

func Parse(filename string) *day01 {
	var (
		data   = utils.ReadLines(filename)
		depths = make([]int, len(data))
	)

	for i, line := range data {
		depths[i] = utils.Atoi(line)
	}

	return &day01{depths: depths}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of measurements larger than previous:", day.Part1())
	fmt.Println(
		"ANSWER2: number of sums larger than previous for 3-measurement sliding window:",
		day.Part2(),
	)
}
