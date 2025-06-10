package day06

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type series [26]bool

type day06 struct {
	lines []string
}

func (s *series) lineToSeries(line string) {
	for i := range s {
		s[i] = false
	}
	for _, char := range line {
		s[char-'a'] = true
	}
}

func (s *series) count() (count int) {
	for _, v := range s {
		if v {
			count++
		}
	}
	return
}

func (d *day06) Part1() (total int) {
	var (
		currentGroup series
		unset        = true
	)

	for _, line := range d.lines {
		if line == "" {
			total += currentGroup.count()
			unset = true
		} else {
			if unset {
				unset = false
				currentGroup.lineToSeries(line)
			} else {
				for _, char := range line {
					if char >= 'a' && char <= 'z' {
						currentGroup[char-'a'] = true
					}
				}
			}
		}
	}

	if !unset {
		total += currentGroup.count()
	}

	return
}

func (d *day06) Part2() (total int) {
	var (
		currentGroup, tempGroup series
		unset                   = true
	)

	for _, line := range d.lines {
		if line == "" {
			total += currentGroup.count()
			unset = true
		} else {
			if unset {
				unset = false
				currentGroup.lineToSeries(line)
			} else {
				tempGroup.lineToSeries(line)
				for i := range tempGroup {
					if !tempGroup[i] {
						currentGroup[i] = false
					}
				}
			}
		}
	}

	if !unset {
		total += currentGroup.count()
	}

	return
}

func Parse(filename string) *day06 {
	lines := utils.ReadLines(filename)

	return &day06{lines: lines}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of counts for `yes`:", day.Part1())
	fmt.Println("ANSWER2: sum of counts for `yes` in all members:", day.Part2())
}
