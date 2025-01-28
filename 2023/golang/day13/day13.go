package day13

import (
	"fmt"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type pattern struct {
	rowMasks []int
	colMasks []int
}

type day13 struct {
	patterns []pattern
}

func findReflection(masks []int) int {
	for i := 1; i < len(masks); i++ {
		if isReflection(masks, i-1, i) {
			return i
		}
	}

	return -1
}

func findNextReflection(masks []int, size int) int {
	for i := 1; i < len(masks); i++ {
		if !isReflection(masks, i-1, i) && canBeReflection(masks, i-1, i, size) {
			return i
		}
	}

	return -1
}

func isReflection(masks []int, left, right int) bool {
	for left >= 0 && right < len(masks) &&
		masks[left] == masks[right] {
		left--
		right++
	}

	return left < 0 || right == len(masks)
}

func canBeReflection(masks []int, left, right int, size int) bool {
	for (left >= 0 && right < len(masks)) &&
		(masks[left] == masks[right] ||
			differsByOne(masks[left], masks[right], size)) {
		left--
		right++
	}

	return left < 0 || right == len(masks)
}

func differsByOne(a, b int, size int) bool {
	if a == b {
		return false
	}

	if a > b {
		a, b = b, a
	}

	for i := range size {
		if (1<<i)|a == b {
			return true
		}
	}

	return false
}

func parsePattern(block []string) pattern {
	rowMasks := make([]int, len(block))
	colMasks := make([]int, len(block[0]))

	for y, row := range block {
		for x, char := range row {
			if char == '#' {
				rowMasks[y] |= 1 << x
				colMasks[x] |= 1 << y
			}
		}
	}

	return pattern{rowMasks, colMasks}
}

// Borrowed from
// https://github.com/aptcode0/adventofcode/blob/main/2023/day13/solution
func (d *day13) Part1() (sum int) {
	for _, p := range d.patterns {
		rowReflection := findReflection(p.rowMasks)
		if rowReflection != -1 {
			sum += (rowReflection) * 100
		} else {
			colReflection := findReflection(p.colMasks)
			sum += colReflection
		}
	}

	return
}

func (d *day13) Part2() (sum int) {
	for _, p := range d.patterns {
		rowReflection := findNextReflection(p.rowMasks, len(p.colMasks))
		if rowReflection != -1 {
			sum += (rowReflection) * 100
		} else {
			colReflection := findNextReflection(p.colMasks, len(p.rowMasks))
			sum += colReflection
		}
	}
	return
}

func Parse(filename string) *day13 {
	var (
		data     = utils.ReadLines(filename)
		patterns []pattern
		block    []string
	)

	for _, line := range data {
		if line == "" {
			patterns = append(patterns, parsePattern(block))
			block = block[:0]
			continue
		}

		block = append(block, line)
	}
	patterns = append(patterns, parsePattern(block))

	return &day13{patterns}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: notes summary:", day.Part1())
	fmt.Println("ANSWER2: notes summary after finding smudge:", day.Part2())
}
