package day06

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day06 struct {
	lanternfishes []int
}

func (d *day06) Part1And2() (part1, part2 int) {
	// each fish's timer can only be 0-8
	totalFishes := make([]int, 9)

	countTotalFishes := func() (total int) {
		for _, f := range totalFishes {
			total += f
		}
		return
	}

	for _, f := range d.lanternfishes {
		totalFishes[f]++
	}

	for day := range 256 {
		fishesToSpawn := totalFishes[0]
		// Decrease days left for fishes
		for i := 0; i < 8; i++ {
			totalFishes[i] = totalFishes[i+1]
		}
		// New fishes start with 8 days
		totalFishes[8] = fishesToSpawn
		// Fishes with 0 days left reset to 6 days
		totalFishes[6] += fishesToSpawn

		if day == 79 {
			part1 = countTotalFishes()
		}
	}

	part2 = countTotalFishes()

	return
}

func Parse(filename string) *day06 {
	var (
		data          = utils.ReadTrimmed(filename)
		split         = strings.Split(data, ",")
		lanternfishes = make([]int, len(split))
	)

	for i, s := range split {
		lanternfishes[i] = utils.Atoi(s)
	}
	return &day06{lanternfishes: lanternfishes}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println("ANSWER1: number of lanternfish after 80 days:", part1)
	fmt.Println("ANSWER2: number of lanternfish after 256 days:", part2)
}
