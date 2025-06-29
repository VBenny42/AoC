package day09

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day09 struct {
	preamble int
	numbers  []int
	part1    int
}

func (d *day09) isValid(index int) bool {
	for i := index - d.preamble; i < index; i++ {
		for j := i + 1; j < index; j++ {
			if d.numbers[i]+d.numbers[j] == d.numbers[index] {
				return true
			}
		}
	}
	return false
}

func (d *day09) Part1() int {
	for i := d.preamble; i < len(d.numbers); i++ {
		if !d.isValid(i) {
			d.part1 = d.numbers[i]
			return d.numbers[i]
		}
	}

	return -1
}

func (d *day09) Part2() int {
	for i := 0; i < len(d.numbers); i++ {
		sum := 0
		min, max := d.numbers[i], d.numbers[i]
		for j := i; j < len(d.numbers); j++ {
			sum += d.numbers[j]
			if d.numbers[j] < min {
				min = d.numbers[j]
			}
			if d.numbers[j] > max {
				max = d.numbers[j]
			}
			if sum == d.part1 {
				return min + max
			} else if sum > d.part1 {
				break
			}
		}
	}

	return -1
}

func Parse(filename string, preamble int) *day09 {
	var (
		lines   = utils.ReadLines(filename)
		numbers = make([]int, len(lines))
		num     int
	)

	for i, line := range lines {
		fmt.Sscanf(line, "%d", &num)
		numbers[i] = num
	}
	return &day09{preamble: preamble, numbers: numbers}
}

func Solve(filename string) {
	day := Parse(filename, 25)

	fmt.Println("ANSWER1: first number that does not follow property:", day.Part1())
	fmt.Println("ANSWER2: encryption weakness:", day.Part2())
}
