package day04

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type pair struct {
	first  [2]int
	second [2]int
}

type day04 struct {
	pairs []pair
}

func (d *day04) Part1() int {
	total := 0

	for _, p := range d.pairs {
		if ((p.first[0] <= p.second[0]) && (p.first[1] >= p.second[1])) ||
			((p.first[0] >= p.second[0]) && (p.first[1] <= p.second[1])) {
			total++
		}
	}

	return total
}

func (d *day04) Part2() int {
	total := 0

	for _, p := range d.pairs {
		if ((p.first[0] <= p.second[0]) && (p.first[1] >= p.second[0])) ||
			((p.first[0] <= p.second[1]) && (p.first[1] >= p.second[1])) ||
			((p.first[0] >= p.second[0]) && (p.first[1] <= p.second[1])) {
			total++
		}
	}

	return total
}

func Parse(filename string) *day04 {
	data := utils.ReadLines(filename)

	pairs := make([]pair, len(data))

	for i, line := range data {
		fmt.Sscanf(line, "%d-%d,%d-%d",
			&pairs[i].first[0],
			&pairs[i].first[1],
			&pairs[i].second[0],
			&pairs[i].second[1])
	}

	return &day04{pairs}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: pairs with a proper subset:", day.Part1())
	fmt.Println("ANSWER2: pairs with a subset:", day.Part2())
}
