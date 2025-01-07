package day02

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type strategy struct {
	opponent rune
	choice   rune
}

type day02 struct {
	strategies []strategy
}

const (
	loss = 0
	draw = 3
	win  = 6
)

func (d *day02) Part1() int {
	total := 0

	outcomes := map[rune]map[rune]int{
		'A': {'X': 4, 'Y': 8, 'Z': 3},
		'B': {'X': 1, 'Y': 5, 'Z': 9},
		'C': {'X': 7, 'Y': 2, 'Z': 6},
	}

	for _, s := range d.strategies {
		total += outcomes[s.opponent][s.choice]
	}

	return total
}

func (d *day02) Part2() int {
	total := 0

	outcomes := map[rune]map[rune]int{
		'A': {'X': 3, 'Y': 4, 'Z': 8},
		'B': {'X': 1, 'Y': 5, 'Z': 9},
		'C': {'X': 2, 'Y': 6, 'Z': 7},
	}

	for _, s := range d.strategies {
		total += outcomes[s.opponent][s.choice]
	}

	return total
}

func Parse(filename string) *day02 {
	data := utils.ReadLines(filename)

	strategies := make([]strategy, len(data))

	for i, line := range data {
		strategies[i].opponent = rune(line[0])
		strategies[i].choice = rune(line[2])
	}

	return &day02{strategies}
}

func Solve(filename string) {
	d := Parse(filename)

	fmt.Println("ANSWER1: total score:", d.Part1())
	fmt.Println("ANSWER2: total score with end:", d.Part2())
}
