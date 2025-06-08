package day01

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day01 struct {
}

func Parse(filename string) *day01 {
	return &day01{}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: <answer1>:", day.Part1())
}
