package day09

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type sequence []int

type day09 struct {
	sequences []sequence
}

func generateBelow(sequence sequence) (sequence, bool) {
	below := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		var (
			curr = sequence[i]
			next = sequence[i+1]
		)
		below[i] = next - curr
	}

	allZero := true
	for _, v := range below {
		if v != 0 {
			allZero = false
			break
		}
	}

	return below, allZero
}

func extrapolateLast(sequence sequence) int {
	below, allZero := generateBelow(sequence)

	if allZero {
		return sequence[len(sequence)-1]
	}

	return sequence[len(sequence)-1] + extrapolateLast(below)
}

func extrapolateFirst(sequence sequence) int {
	below, allZero := generateBelow(sequence)

	if allZero {
		return sequence[0]
	}

	return sequence[0] - extrapolateFirst(below)
}

func (d *day09) Part1() (sum int) {
	for _, sequence := range d.sequences {
		sum += extrapolateLast(sequence)
	}

	return
}

func (d *day09) Part2() (sum int) {
	for _, sequence := range d.sequences {
		sum += extrapolateFirst(sequence)
	}

	return
}

func Parse(filename string) *day09 {
	var (
		data      = utils.ReadLines(filename)
		sequences = make([]sequence, len(data))
	)

	for i, line := range data {
		fields := strings.Fields(line)
		sequences[i] = make(sequence, len(fields))
		for j, field := range fields {
			sequences[i][j] = utils.Atoi(field)
		}
	}

	return &day09{sequences}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of extrapolated last numbers:", day.Part1())
	fmt.Println("ANSWER2: sum of extrapolated first numbers:", day.Part2())
}
