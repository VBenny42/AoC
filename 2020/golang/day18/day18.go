package day18

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day18 struct {
	input []string
}

func calculate(s string, runningTotal int) int {
	var operator byte = 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		case '+':
			operator = '+'
		case '*':
			operator = '*'
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			value := int(s[i] - '0')
			if operator == 0 {
				runningTotal = value
			} else if operator == '+' {
				runningTotal += value
			} else if operator == '*' {
				runningTotal *= value
			}
			operator = 0
		case '(':
			parenCount := 1
			j := i + 1
			for j < len(s) && parenCount > 0 {
				if s[j] == '(' {
					parenCount++
				} else if s[j] == ')' {
					parenCount--
				}
				j++
			}
			if parenCount == 0 {

				value := calculate(s[i+1:j-1], 0)
				if operator == 0 {
					runningTotal = value
				} else if operator == '+' {
					runningTotal += value
				} else if operator == '*' {
					runningTotal *= value
				}
				operator = 0
				i = j - 1
			}
		case ')':
			continue
		default:
			panic(fmt.Sprintf("Unexpected character '%c' in input", s[i]))
		}
	}
	return runningTotal
}

// Claude did this lol
func calculate2(s string) int {
	// First, handle all parentheses by recursively evaluating them
	for {
		start := -1
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				start = i
			} else if s[i] == ')' {
				if start != -1 {
					// Replace the parentheses expression with its calculated value
					value := calculate2(s[start+1 : i])
					s = s[:start] + fmt.Sprintf("%d", value) + s[i+1:]
					break
				}
			}
		}
		if start == -1 {
			break // No more parentheses
		}
	}

	// Now handle addition first (higher precedence than multiplication)
	for {
		found := false
		for i := 0; i < len(s); i++ {
			if s[i] == '+' {
				// Find the operands on both sides
				left := ""
				right := ""

				// Get left operand
				j := i - 1
				for j >= 0 && s[j] == ' ' {
					j--
				}
				for j >= 0 && (s[j] >= '0' && s[j] <= '9') {
					left = string(s[j]) + left
					j--
				}
				leftStart := j + 1

				// Get right operand
				j = i + 1
				for j < len(s) && s[j] == ' ' {
					j++
				}
				for j < len(s) && (s[j] >= '0' && s[j] <= '9') {
					right = right + string(s[j])
					j++
				}
				rightEnd := j

				if left != "" && right != "" {
					leftVal := 0
					rightVal := 0

					// Convert strings to integers
					for _, ch := range left {
						leftVal = leftVal*10 + int(ch-'0')
					}
					for _, ch := range right {
						rightVal = rightVal*10 + int(ch-'0')
					}

					result := leftVal + rightVal
					s = s[:leftStart] + fmt.Sprintf("%d", result) + s[rightEnd:]
					found = true
					break
				}
			}
		}
		if !found {
			break // No more additions
		}
	}

	// Finally, handle multiplication (left to right)
	result := 0
	operator := byte(0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		case '*':
			operator = '*'
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// Parse the full number
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			i-- // Adjust for the loop increment

			if operator == 0 {
				result = num
			} else if operator == '*' {
				result *= num
			}
			operator = 0
		}
	}

	return result
}

func (d *day18) Part1() (sum int) {
	for _, line := range d.input {
		sum += calculate(line, 0)
	}

	return
}

func (d *day18) Part2() (sum int) {
	for _, line := range d.input {
		sum += calculate2(line)
	}

	return
}

func Parse(filename string) *day18 {
	return &day18{input: utils.ReadLines(filename)}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of resulting values:", day.Part1())
	fmt.Println("ANSWER2: sum of resulting values with addition first:", day.Part2())
}
