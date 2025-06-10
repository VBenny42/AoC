package day05

import (
	"fmt"
	"slices"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day05 struct {
	lines []string
	ids   []int
}

func calculateSeatID(seat string) int {
	row := 0
	col := 0

	for i := 0; i < 7; i++ {
		row <<= 1
		if seat[i] == 'B' {
			row |= 1
		}
	}

	for i := 7; i < 10; i++ {
		col <<= 1
		if seat[i] == 'R' {
			col |= 1
		}
	}

	return row*8 + col
}

func (d *day05) Part1() (maxSeatID int) {
	for i, line := range d.lines {
		seatID := calculateSeatID(line)
		d.ids[i] = seatID
		maxSeatID = max(maxSeatID, seatID)
	}

	return
}

func (d *day05) Part2() (missingSeatID int) {
	slices.Sort(d.ids)

	for i := 0; i < len(d.ids)-1; i++ {
		if d.ids[i+1] != d.ids[i]+1 {
			missingSeatID = d.ids[i] + 1
			break
		}
	}

	return
}

func Parse(filename string) *day05 {
	lines := utils.ReadLines(filename)
	return &day05{lines: lines, ids: make([]int, len(lines))}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: highest seat ID on a boarding pass:", day.Part1())
	fmt.Println("ANSWER2: missing seat ID:", day.Part2())
}
