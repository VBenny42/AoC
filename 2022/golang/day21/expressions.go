package day21

import (
	"fmt"
	"strconv"
	"strings"
)

func evaluateEquation(equation string) (int, error) {
	// Remove all spaces
	equation = strings.ReplaceAll(equation, " ", "")

	// Base case: single number
	if num, err := strconv.Atoi(equation); err == nil {
		return num, nil
	}

	// Find matching parentheses
	if equation[0] != '(' || equation[len(equation)-1] != ')' {
		return 0, fmt.Errorf("invalid equation format: %s", equation)
	}

	// Remove outer parentheses
	equation = equation[1 : len(equation)-1]

	// Find the main operator (the one not inside parentheses)
	parenthesesCount := 0
	var operator rune
	var operatorIndex int

	for i, char := range equation {
		switch char {
		case '(':
			parenthesesCount++
		case ')':
			parenthesesCount--
		case '+', '-', '*', '/':
			if parenthesesCount == 0 {
				operator = char
				operatorIndex = i
			}
		}
	}

	if operatorIndex == 0 {
		return 0, fmt.Errorf("no valid operator found in equation: %s", equation)
	}

	// Split into left and right parts
	left := equation[:operatorIndex]
	right := equation[operatorIndex+1:]

	// Recursively evaluate both parts
	leftValue, err := evaluateEquation(left)
	if err != nil {
		return 0, err
	}

	rightValue, err := evaluateEquation(right)
	if err != nil {
		return 0, err
	}

	// Apply the operator
	if evaluator, ok := operators[operator]; ok {
		return evaluator(leftValue, rightValue), nil
	}

	return 0, fmt.Errorf("unknown operator: %c", operator)
}
