package day04

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type card struct {
	winMask [2]uint64
	mine    []int
}

const (
	low  = 0
	high = 1
)

type day04 struct {
	cards []card
}

func inMask(mask [2]uint64, val int) bool {
	if val < 64 {
		return mask[low]&(1<<val) != 0
	}

	return mask[high]&(1<<(val-64)) != 0
}

func (d *day04) Part1() (sum int) {
	for _, c := range d.cards {
		var points int

		for _, m := range c.mine {
			if inMask(c.winMask, m) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		sum += points
	}

	return
}

func (d *day04) Part2() (sum int) {
	instances := make(map[int]int)

	for i := range d.cards {
		instances[i] = 1
	}

	for i, c := range d.cards {
		var matches int

		for _, m := range c.mine {
			if inMask(c.winMask, m) {
				matches++
			}
		}

		for j := 1; j <= matches; j++ {
			instances[i+j] += instances[i]
		}
	}

	for _, v := range instances {
		sum += v
	}

	return
}

func Parse(filename string) *day04 {
	var (
		data  = utils.ReadLines(filename)
		cards = make([]card, len(data))
	)

	for i, line := range data {
		parts := strings.Split(
			strings.Split(line, ": ")[1],
			"|")

		var (
			winning = strings.Fields(parts[0])
			mine    = strings.Fields(parts[1])
		)

		for _, w := range winning {
			val := utils.Atoi(w)
			if val < 64 {
				cards[i].winMask[low] |= 1 << val
			} else {
				cards[i].winMask[high] |= 1 << (val - 64)
			}
		}

		cards[i].mine = make([]int, len(mine))
		for j, m := range mine {
			cards[i].mine[j] = utils.Atoi(m)
		}
	}

	return &day04{cards}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: total points:", day.Part1())
	fmt.Println("ANSWER2: total scratchcards:", day.Part2())
}
