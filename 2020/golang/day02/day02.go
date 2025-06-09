package day02

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type line struct {
	letter     rune
	pos1, pos2 int
	password   string
}

type day02 struct {
	lines []line
}

func (l *line) isValid1() bool {
	count := 0
	for _, c := range l.password {
		if c == l.letter {
			count++
		}
	}
	return count >= l.pos1 && count <= l.pos2
}

func (l *line) isValid2() bool {
	// if l.pos1 < 1 || l.pos2 < 1 || l.pos1 > len(l.password) || l.pos2 > len(l.password) {
	// 	return false
	// }
	first := rune(l.password[l.pos1-1]) == l.letter
	second := rune(l.password[l.pos2-1]) == l.letter
	return (first || second) && !(first && second)
}

func (d *day02) Part1And2() (validCount1, validCount2 int) {
	for _, l := range d.lines {
		if l.isValid1() {
			validCount1++
		}
		if l.isValid2() {
			validCount2++
		}
	}
	return
}

func Parse(filename string) *day02 {
	var (
		day        day02
		pos1, pos2 int
		letter     rune
		password   string
	)

	for _, l := range utils.ReadLines(filename) {
		_, err := fmt.Sscanf(l, "%d-%d %c: %s", &pos1, &pos2, &letter, &password)
		if err != nil {
			panic(err)
		}
		day.lines = append(day.lines, line{
			pos1:     pos1,
			pos2:     pos2,
			letter:   letter,
			password: password,
		})
	}
	return &day
}

func Solve(filename string) {
	day := Parse(filename)

	validCount1, validCount2 := day.Part1And2()

	fmt.Println("ANSWER1: number of valid passwords:", validCount1)
	fmt.Println(
		"ANSWER2: number of valid passwords according to new interpretation:",
		validCount2,
	)
}
