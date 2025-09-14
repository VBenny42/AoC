package day04

import (
	"fmt"

	"github.com/VBenny42/AoC/2019/golang/utils"
)

type day04 struct {
	start, end int
}

type password [6]int

func (p *password) isValid() (valid bool) {
	for i := 1; i < 6; i++ {
		if p[i] == p[i-1] {
			valid = true
		}
	}
	return
}

func (p *password) isValid2() (valid bool) {
	var count int
	for i := 1; i < 6; i++ {
		if p[i] == p[i-1] {
			count++
		} else {
			if count == 1 {
				valid = true
			}
			count = 0
		}
	}
	if count == 1 {
		valid = true
	}

	return
}

func (day *day04) Part1And2() (count1, count2 int) {
	for a := 1; a <= 9; a++ {
		for b := a; b <= 9; b++ {
			for c := b; c <= 9; c++ {
				for d := c; d <= 9; d++ {
					for e := d; e <= 9; e++ {
						for f := e; f <= 9; f++ {
							num := a*100000 + b*10000 + c*1000 + d*100 + e*10 + f
							if num < day.start || num > day.end {
								continue
							}
							p := password{a, b, c, d, e, f}
							if p.isValid() {
								count1++

								// Valid part 2 password implies valid part 1 password
								if p.isValid2() {
									count2++
								}
							}
						}
					}
				}
			}
		}
	}

	return
}

func Parse(filename string) *day04 {
	var (
		day  day04
		line = utils.ReadTrimmed(filename)
	)

	fmt.Sscanf(line, "%d-%d", &day.start, &day.end)

	return &day
}

func Solve(filename string) {
	day := Parse(filename)

	count1, count2 := day.Part1And2()

	fmt.Println("ANSWER1: number of valid passwords:", count1)
	fmt.Println("ANSWER2: number of valid passwords with new double constraint:", count2)
}
