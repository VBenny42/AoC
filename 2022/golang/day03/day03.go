package day03

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
	"github.com/deckarep/golang-set/v2"
)

type day03 struct {
	rucksacks []string
}

func getPriority(r rune) int {
	priority := int(r) - 'a' + 1
	if r <= 'Z' {
		priority += 58
	}
	return priority
}

func (d *day03) Part1() int {
	total := 0

	for _, rucksack := range d.rucksacks {
		common, ok := mapset.NewSet([]rune(rucksack[:len(rucksack)/2])...).
			Intersect(mapset.NewSet([]rune(rucksack[len(rucksack)/2:])...)).
			Pop()
		if !ok {
			panic("no common items found")
		}
		total += getPriority(common)
	}

	return total
}

func (d *day03) Part2() int {
	total := 0

	for i := 0; i < len(d.rucksacks); i += 3 {
		common, ok := mapset.NewSet([]rune(d.rucksacks[i])...).
			Intersect(mapset.NewSet([]rune(d.rucksacks[i+1])...)).
			Intersect(mapset.NewSet([]rune(d.rucksacks[i+2])...)).
			Pop()
		if !ok {
			panic("no common items found")
		}
		total += getPriority(common)
	}

	return total
}

func Parse(filename string) *day03 {
	data := utils.ReadLines(filename)

	return &day03{rucksacks: data}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of priorities of common items:", day.Part1())
	fmt.Println("ANSWER2: sum of priorities of common items in groups of 3:", day.Part2())
}
