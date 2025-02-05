package day02

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type (
	direction int
	command   struct {
		dir    direction
		amount int
	}
)

type day02 struct {
	commands []command
}

const (
	forward direction = iota
	down
	up
)

func (d *day02) Part1() int {
	var horizontal, depth int

	for _, c := range d.commands {
		switch c.dir {
		case forward:
			horizontal += c.amount
		case down:
			depth += c.amount
		case up:
			depth -= c.amount
		}
	}

	return horizontal * depth
}

func (d *day02) Part2() int {
	var horizontal, depth, aim int

	for _, c := range d.commands {
		switch c.dir {
		case forward:
			horizontal += c.amount
			depth += c.amount * aim
		case down:
			aim += c.amount
		case up:
			aim -= c.amount
		}
	}

	return horizontal * depth
}

func Parse(filename string) *day02 {
	var (
		data     = utils.ReadLines(filename)
		commands = make([]command, len(data))
	)

	for i, line := range data {
		split := strings.SplitN(line, " ", 2)
		switch split[0] {
		case "forward":
			commands[i].dir = forward
		case "down":
			commands[i].dir = down
		case "up":
			commands[i].dir = up
		}
		commands[i].amount = utils.Atoi(split[1])
	}

	return &day02{commands}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: product of final horizontal position and depth:", day.Part1())
	fmt.Println(
		"ANSWER2: product of final horizontal position and depth with aim:",
		day.Part2(),
	)
}
