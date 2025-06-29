package day15

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day15 struct {
	startingNumbers []int
}

func (d *day15) traverse(target int) int {
	var (
		lastSpoken = make([]int, target)
		lastNumber = d.startingNumbers[len(d.startingNumbers)-1]
		nextNumber int
	)

	for i, num := range d.startingNumbers[:len(d.startingNumbers)-1] {
		lastSpoken[num] = i + 1
	}

	for turn := len(d.startingNumbers); turn < target; turn++ {
		if previousTurn := lastSpoken[lastNumber]; previousTurn != 0 {
			nextNumber = turn - previousTurn
		} else {
			nextNumber = 0
		}

		lastSpoken[lastNumber] = turn

		lastNumber = nextNumber
	}

	return lastNumber
}

func (d *day15) Part1() int {
	return d.traverse(2020)
}

func (d *day15) Part2() int {
	return d.traverse(30000000)
}

func Parse(filename string) *day15 {
	line := utils.ReadTrimmed(filename)
	var startingNumbers []int

	for _, numStr := range strings.Split(line, ",") {
		startingNumbers = append(startingNumbers, utils.Atoi(numStr))
	}

	return &day15{startingNumbers: startingNumbers}
}

func Solve(filename string) {
	day := Parse(filename)
	fmt.Println("ANSWER1: 2020th number spoken:", day.Part1())
	fmt.Println("ANSWER2: 30000000th number spoken:", day.Part2())
}
