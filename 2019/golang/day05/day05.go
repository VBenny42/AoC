package day05

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2019/golang/utils"
)

type day05 struct {
	program []int
}

func processDay05(program []int, input int) (outputs []int) {
	var ip int

	for ip < len(program) {
		instruction := program[ip]
		opcode := instruction % 100

		mode1 := (instruction / 100) % 10
		mode2 := (instruction / 1000) % 10

		if opcode == 99 {
			break
		}

		getValue := func(param, mode int) int {
			if mode == 0 { // position mode
				return program[param]
			} else { // immediate mode
				return param
			}
		}

		switch opcode {
		case 1: // addition
			pos1 := program[ip+1]
			pos2 := program[ip+2]
			pos3 := program[ip+3]
			val1 := getValue(pos1, mode1)
			val2 := getValue(pos2, mode2)
			program[pos3] = val1 + val2
			ip += 4
		case 2: // multiplication
			pos1 := program[ip+1]
			pos2 := program[ip+2]
			pos3 := program[ip+3]
			val1 := getValue(pos1, mode1)
			val2 := getValue(pos2, mode2)
			program[pos3] = val1 * val2
			ip += 4
		case 3: // input
			pos1 := program[ip+1]
			program[pos1] = input
			ip += 2
		case 4: // output
			pos1 := program[ip+1]
			val1 := getValue(pos1, mode1)
			outputs = append(outputs, val1)
			ip += 2
		case 5: // jump-if-true
			pos1 := program[ip+1]
			pos2 := program[ip+2]
			val1 := getValue(pos1, mode1)
			val2 := getValue(pos2, mode2)
			if val1 != 0 {
				ip = val2
			} else {
				ip += 3
			}
		case 6: // jump-if-false
			pos1 := program[ip+1]
			pos2 := program[ip+2]
			val1 := getValue(pos1, mode1)
			val2 := getValue(pos2, mode2)
			if val1 == 0 {
				ip = val2
			} else {
				ip += 3
			}
		case 7: // less than
			pos1 := program[ip+1]
			pos2 := program[ip+2]
			pos3 := program[ip+3]
			val1 := getValue(pos1, mode1)
			val2 := getValue(pos2, mode2)
			if val1 < val2 {
				program[pos3] = 1
			} else {
				program[pos3] = 0
			}
			ip += 4
		case 8: // equals
			pos1 := program[ip+1]
			pos2 := program[ip+2]
			pos3 := program[ip+3]
			val1 := getValue(pos1, mode1)
			val2 := getValue(pos2, mode2)
			if val1 == val2 {
				program[pos3] = 1
			} else {
				program[pos3] = 0
			}
			ip += 4
		default:
			panic(fmt.Sprintf("Unknown opcode %d at position %d", opcode, ip))
		}
	}

	return
}

func (d *day05) Part1() int {
	programCopy := make([]int, len(d.program))
	copy(programCopy, d.program)

	// Run the program with input 1 (air conditioner unit ID)
	outputs := processDay05(programCopy, 1)

	// The diagnostic code is the last output
	if len(outputs) > 0 {
		return outputs[len(outputs)-1]
	}

	return 0
}

func (d *day05) Part2(id int) int {
	programCopy := make([]int, len(d.program))
	copy(programCopy, d.program)

	// Run the program with id
	outputs := processDay05(programCopy, id)

	// The diagnostic code is the last output
	if len(outputs) > 0 {
		return outputs[len(outputs)-1]
	}

	return 0
}

func Parse(filename string) *day05 {
	line := utils.ReadTrimmed(filename)
	return &day05{
		program: utils.MapSlices(
			strings.Split(line, ","),
			utils.Atoi,
		),
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: diagnostic code produced:", day.Part1())
	fmt.Println("ANSWER2: diagnostic code for system ID 5:", day.Part2(5))
}
