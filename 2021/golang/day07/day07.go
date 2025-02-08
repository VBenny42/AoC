package day07

import (
	"fmt"
	"sort"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day07 struct {
	crabs []int
}

// Guessing that the spot that all crabs should align to would be the median
// of the crabs. This is because the median would be in the middle of the
// list, so the least amount of movement would be required to get all crabs to that spot.
// Either that or the mode, but I'm going with the median.
func (d *day07) Part1() (sum int) {
	sort.Ints(d.crabs)
	median := d.crabs[len(d.crabs)/2]

	for _, crab := range d.crabs {
		sum += utils.Abs(crab - median)
	}

	return
}

// Need to move to mean now
func (d *day07) Part2() int {
	var total int

	for _, crab := range d.crabs {
		total += crab
	}

	var (
		floor = (total - len(d.crabs)/2) / len(d.crabs)
		ceil  = (total + len(d.crabs)/2) / len(d.crabs)
	)

	distanceToVal := func(val int) (sum int) {
		for _, crab := range d.crabs {
			distance := utils.Abs(crab - val)
			for distance > 0 {
				sum += distance
				distance--
			}
		}
		return
	}

	return min(distanceToVal(floor), distanceToVal(ceil))
}

func Parse(filename string) *day07 {
	var (
		data  = utils.ReadTrimmed(filename)
		split = strings.Split(data, ",")
		crabs = make([]int, len(split))
	)

	for i, s := range split {
		crabs[i] = utils.Atoi(s)
	}

	return &day07{crabs: crabs}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: fuel spent to get to optimal horizontal position:", day.Part1())
	fmt.Println("ANSWER2: fuel spent to get to new optimal horizontal position:", day.Part2())
}
