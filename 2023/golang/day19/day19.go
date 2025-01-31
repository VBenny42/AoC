package day19

import (
	"fmt"
)

type (
	part [4]int

	partIdx   int
	operator  rune
	condition struct {
		idx   partIdx
		op    operator
		value int
	}

	rule struct {
		cond   *condition
		target string
	}
	workflow struct {
		rules []rule
	}
)

const (
	x partIdx = iota
	m
	a
	s
)

const (
	gt operator = '>'
	lt operator = '<'
)

type day19 struct {
	workflows map[string]workflow
	parts     []part
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of accepted parts:", day.Part1())
	fmt.Println(
		"ANSWER2: number of distinct combinations that will be accepted:",
		day.Part2(),
	)
}
