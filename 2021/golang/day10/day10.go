package day10

import (
	"fmt"
	"sort"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type day10 struct {
	lines []string
}

const (
	lRoundBracket  = '('
	rRoundBracket  = ')'
	lSquareBracket = '['
	rSquareBracket = ']'
	lCurlyBracket  = '{'
	rCurlyBracket  = '}'
	lAngleBracket  = '<'
	rAngleBracket  = '>'
)

var errorToPoints = map[rune]int{
	rRoundBracket:  3,
	rSquareBracket: 57,
	rCurlyBracket:  1197,
	rAngleBracket:  25137,
}

var autoCorrectToPoints = map[rune]int{
	// Mapping directly to left brackets instead of finding opposite then mapping
	lRoundBracket:  1,
	lSquareBracket: 2,
	lCurlyBracket:  3,
	lAngleBracket:  4,
}

func getOpposite(l rune) rune {
	switch l {
	case lRoundBracket:
		return rRoundBracket
	case lSquareBracket:
		return rSquareBracket
	case lCurlyBracket:
		return rCurlyBracket
	case lAngleBracket:
		return rAngleBracket
	}
	fmt.Println("unmatched bracket")
	return 0
}

func (d *day10) Part1() (sum int) {
	// Need to implement a stack to keep track of the current depth
	// and current chunk start bracket
	var indicesToRemove []int
	var stack []rune
	for i, line := range d.lines {
		// Change to linked list?
		stack = stack[:0]
		for _, r := range line {
			switch r {
			case lRoundBracket, lSquareBracket, lCurlyBracket, lAngleBracket:
				stack = append(stack, r)
			case rRoundBracket, rSquareBracket, rCurlyBracket, rAngleBracket:
				// pop the last bracket
				lastBracket := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if r != getOpposite(lastBracket) {
					// wrong bracket
					sum += errorToPoints[r]
					indicesToRemove = append(indicesToRemove, i)
					break
				}
			}
		}
	}

	// Remove the lines with syntax errors for part 2
	d.lines = utils.RemoveIndices(d.lines, indicesToRemove...)

	return
}

func (d *day10) Part2() int {
	// Need to add the brackets to make the line complete
	autoCorrectScores := make([]int, len(d.lines))
	var stack []rune
	for i, line := range d.lines {
		// Use stack like before, now add the missing brackets
		stack = stack[:0]
		for _, r := range line {
			switch r {
			case lRoundBracket, lSquareBracket, lCurlyBracket, lAngleBracket:
				stack = append(stack, r)
			case rRoundBracket, rSquareBracket, rCurlyBracket, rAngleBracket:
				// pop the last bracket
				lastBracket := stack[len(stack)-1]
				if r == getOpposite(lastBracket) {
					// Pair is correct, remove pair from stack
					stack = stack[:len(stack)-1]
				}
				// Do nothing if the bracket is wrong, gonna add at the end
			}
		}

		for j := len(stack) - 1; j >= 0; j-- {
			autoCorrectScores[i] *= 5
			autoCorrectScores[i] += autoCorrectToPoints[stack[j]]
		}
	}

	sort.Ints(autoCorrectScores)

	return autoCorrectScores[len(autoCorrectScores)/2]
}

func Parse(filename string) *day10 {
	return &day10{lines: utils.ReadLines(filename)}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: total syntax error score:", day.Part1())
	fmt.Println("ANSWER2: total auto-correct score:", day.Part2())
}
