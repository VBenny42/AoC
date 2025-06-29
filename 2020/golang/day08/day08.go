package day08

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type operation struct {
	op  string
	arg int
}

type day08 struct {
	operations []operation
}

func (d *day08) Part1() (accumuluator int, loopDetected bool) {
	isExecuted := make([]bool, len(d.operations))
	for i := 0; i < len(d.operations); {
		if isExecuted[i] {
			return accumuluator, true // infinite loop detected
		}
		isExecuted[i] = true

		switch d.operations[i].op {
		case "nop":
			i++
		case "acc":
			accumuluator += d.operations[i].arg
			i++
		case "jmp":
			i += d.operations[i].arg
		default:
			panic("unknown operation: " + d.operations[i].op)
		}
	}

	return accumuluator, false
}

func (d *day08) Part2() int {
	for i := 0; i < len(d.operations); i++ {
		// Try to fix the program by changing one operation
		originalOp := d.operations[i].op
		if originalOp == "acc" {
			continue // can't change acc operation
		}

		d.operations[i].op = "jmp"
		if originalOp == "jmp" {
			d.operations[i].op = "nop"
		}

		accumuluator, loopDetected := d.Part1()
		if !loopDetected {
			return accumuluator // fixed program, return accumuluator
		}

		// Restore the original operation
		d.operations[i].op = originalOp
	}

	return -1
}

func Parse(filename string) *day08 {
	lines := utils.ReadLines(filename)
	ops := make([]operation, len(lines))

	for i, line := range lines {
		var op string
		var arg int
		fmt.Sscanf(line, "%s %d", &op, &arg)
		ops[i] = operation{op: op, arg: arg}
	}

	return &day08{operations: ops}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, _ := day.Part1()
	fmt.Println("ANSWER1: accumuluator before infinite loop starts:", part1)
	fmt.Println("ANSWER2: accumuluator after fixing the program:", day.Part2())
}
