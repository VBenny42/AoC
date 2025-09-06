package day02

import (
	"fmt"
	"slices"
	"strings"

	"github.com/VBenny42/AoC/2019/golang/utils"
)

type day02 struct {
	program []int
}

func processDay02(program []int) {
	for i := 0; i < len(program); i += 4 {
		opcode := program[i]
		if opcode == 99 {
			break
		}
		pos1 := program[i+1]
		pos2 := program[i+2]
		pos3 := program[i+3]
		switch opcode {
		case 1:
			program[pos3] = program[pos1] + program[pos2]
		case 2:
			program[pos3] = program[pos1] * program[pos2]
		default:
			panic(fmt.Sprintf("Unknown opcode %d at position %d", opcode, i))
		}
	}
}

func (d *day02) Part1() int {
	// Reset program to initial state
	progCopy := slices.Clone(d.program)
	progCopy[1] = 12
	progCopy[2] = 2
	processDay02(progCopy)
	return progCopy[0]
}

func (d *day02) Part2() int {
	noun, verb := 0, 0
	output := 19690720
	progCopy := slices.Clone(d.program)

	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			// Create a copy of the initial program
			copy(progCopy, d.program)

			// Set noun and verb
			progCopy[1] = noun
			progCopy[2] = verb

			// Process the program
			processDay02(progCopy)

			// Check if the output matches the desired value
			if progCopy[0] == output {
				return 100*noun + verb
			}
		}
	}

	return -1 // Not found
}

func Parse(filename string) *day02 {
	line := utils.ReadTrimmed(filename)
	return &day02{
		program: utils.MapSlices(
			strings.Split(line, ","),
			utils.Atoi,
		),
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: value at position 0 after halt:", day.Part1())
	fmt.Println("ANSWER2: 100 * noun + verb to produce 19690720:", day.Part2())
}
