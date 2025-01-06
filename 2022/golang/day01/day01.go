package day01

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type day01 struct {
	totalCalories []int
}

func (d *day01) Part1() int {
	return d.totalCalories[0]
}

func (d *day01) Part2() int {
	return d.totalCalories[0] + d.totalCalories[1] + d.totalCalories[2]
}

func Parse(filename string) *day01 {
	data := utils.SplitLines(filename)

	calories := make([]int, 0)

	current := 0

	for _, line := range data {
		if line == "" {
			calories = append(calories, current)
			current = 0
			continue
		}

		calorie, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		current += calorie
	}
	calories = append(calories, current)
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	return &day01{calories}
}

func Solve(filename string) {
	d := Parse(filename)

	fmt.Println("ANSWER1: highest carried calories:", d.Part1())
	fmt.Println("ANSWER2: sum of the three highest carried calories:", d.Part2())
}
