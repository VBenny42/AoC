package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type instruction struct {
	opcode bool
	value  int
}

type day10 struct {
	x                 int
	cycleCount        int
	cycleChecks       []int
	indexCheck        int
	signalStrengthSum int
	instructions      []instruction
	builder           strings.Builder
}

func (d *day10) runCycle() {
	d.cycleCount++
	if d.cycleCount == d.cycleChecks[d.indexCheck] {
		d.signalStrengthSum += d.x * d.cycleCount
		d.indexCheck++
	}
}

const (
	rowSize = 40
	padNum  = 3
)

func (d *day10) drawCycle() {
	if utils.Abs((d.cycleCount%rowSize)-d.x) <= 1 {
		d.builder.WriteString("#")
	} else {
		d.builder.WriteString(".")
	}
	d.cycleCount++
	if d.cycleCount%rowSize == 0 {
		d.builder.WriteString("\n")
	}
}

func (d *day10) Part1() int {
	for _, instruction := range d.instructions {
		// noop case
		if instruction.opcode {
			d.runCycle()
		} else {
			// addx case
			for range 2 {
				d.runCycle()
			}
			d.x += instruction.value
		}
		if d.indexCheck == len(d.cycleChecks) {
			break
		}
	}
	return d.signalStrengthSum
}

func (d *day10) Part2() string {
	d.builder = strings.Builder{}

	for _, instruction := range d.instructions {
		if instruction.opcode {
			d.drawCycle()
		} else {
			for range 2 {
				d.drawCycle()
			}
			d.x += instruction.value
		}
	}

	return d.builder.String()
}

func Parse(filename string) *day10 {
	data := utils.ReadLines(filename)

	instructions := make([]instruction, len(data))

	for i, line := range data {
		split := strings.Split(line, " ")
		instructions[i].opcode = len(split) == 1
		if len(split) == 2 {
			instructions[i].value = utils.Must(strconv.Atoi(split[1]))
		}
	}

	return &day10{
		x:                 1,
		cycleCount:        0,
		cycleChecks:       []int{20, 60, 100, 140, 180, 220},
		instructions:      instructions,
		indexCheck:        0,
		signalStrengthSum: 0,
	}
}

func Solve(filename string) {
	fmt.Println("ANSWER1: signal strength sum:", Parse(filename).Part1())
	Parse(filename).Part2()
	// fmt.Printf("ANSWER2: printed to console\n%s", Parse(filename).Part2())
	fmt.Println("ANSWER2: not printing answer, but it's BPJAZGAP")
}
