package day01

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
	"github.com/VBenny42/AoC/2020/golang/utils/set"
)

type day01 struct {
	entries set.Set[int]
}

func (d *day01) Part1() int {
	for _, entry := range d.entries.Values() {
		diff := (2020 - entry)
		if diff < 0 {
			continue
		}
		if d.entries.Contains(diff) {
			return entry * diff
		}
	}
	// Should never happen
	panic("no entries found that sum to 2020")
}

func (d *day01) Part2() int {
	values := d.entries.Values()
	for _, entry1 := range values {
		for _, entry2 := range values {
			if entry1 == entry2 {
				continue
			}
			diff := (2020 - entry1 - entry2)
			if diff < 0 {
				continue
			}
			if d.entries.Contains(diff) {
				return entry1 * entry2 * diff
			}
		}
	}
	// Should never happen
	panic("no entries found that sum to 2020")
}

func Parse(filename string) *day01 {
	var (
		entry int
		day   day01
	)

	day.entries = set.NewSet[int]()

	for _, line := range utils.ReadLines(filename) {
		entry = utils.Atoi(line)
		day.entries.Add(entry)
	}

	return &day
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: product of 2 entries that sum to 2020:", day.Part1())
	fmt.Println("ANSWER2: product of 3 entries that sum to 2020:", day.Part2())
}
