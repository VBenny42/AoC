package day14

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day14 struct {
	template string
	rules    map[string]rune
}

func leastMostDifference(charCount map[rune]int) int {
	var least, most int
	for _, count := range charCount {
		if least == 0 {
			least = count
		}

		least = min(least, count)
		most = max(most, count)
	}

	return most - least
}

func (d *day14) Part1And2() (part1, part2 int) {
	// Initialize pair counts from the template
	pairCounts := make(map[string]int)
	for i := 0; i < len(d.template)-1; i++ {
		pair := d.template[i : i+2]
		pairCounts[pair]++
	}

	// Copy initial character counts
	charCounts := make(map[rune]int)
	for _, char := range d.template {
		charCounts[char]++
	}

	for step := 0; step < 40; step++ {
		newPairCounts := make(map[string]int)

		for pair, count := range pairCounts {
			inBetween, exists := d.rules[pair]
			if !exists {
				newPairCounts[pair] += count
				continue
			}

			// Split into left and right pairs
			var (
				leftPair  = string(pair[0]) + string(inBetween)
				rightPair = string(inBetween) + string(pair[1])
			)

			newPairCounts[leftPair] += count
			newPairCounts[rightPair] += count

			// Add the inserted character count
			charCounts[inBetween] += count
		}

		pairCounts = newPairCounts

		if step == 9 {
			part1 = leastMostDifference(charCounts)
		}
	}

	part2 = leastMostDifference(charCounts)

	return
}

func Parse(filename string) *day14 {
	var (
		data     = utils.ReadLines(filename)
		rules    = make(map[string]rune, len(data[2:]))
		template = data[0]
	)

	for _, line := range data[2:] {
		left, right, found := strings.Cut(line, " -> ")
		if !found {
			panic("invalid input")
		}
		rules[left] = rune(right[0])
	}

	return &day14{
		template: template,
		rules:    rules,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println(
		"ANSWER1: most common element - least common element after 10 steps:",
		part1,
	)
	fmt.Println(
		"ANSWER2: most common element - least common element after 40 steps:",
		part2,
	)
}
