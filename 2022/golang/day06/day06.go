package day06

import (
	"fmt"

	"github.com/VBenny42/AoC/2022/golang/utils"
	"github.com/deckarep/golang-set/v2"
)

type day06 struct {
	message string
}

func allDistinct(window []rune) bool {
	set := mapset.NewSet(rune(window[0]))
	for _, r := range window[1:] {
		if !set.Add(r) {
			return false
		}
	}
	return true
}

func (d *day06) findMarker(distinctAmount int) int {
	for i := distinctAmount; i < len(d.message); i++ {
		if allDistinct([]rune(d.message[i-distinctAmount : i])) {
			return i
		}
	}
	return -1
}

func (d *day06) Part1() int {
	return d.findMarker(4)
}

func (d *day06) Part2() int {
	return d.findMarker(14)
}

func Parse(filename string) *day06 {
	return &day06{
		message: utils.ReadTrimmed(filename),
	}
}

func Solve(filename string) {
	day := Parse(filename)
	fmt.Println("ANSWER1: positions before first marker:", day.Part1())
	fmt.Println("ANSWER2: positions before first marker:", day.Part2())
}
