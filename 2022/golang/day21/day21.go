package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type operation struct {
	left     string
	right    string
	operator rune
}

var operators = map[rune]func(int, int) int{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
	'*': func(a, b int) int { return a * b },
	'/': func(a, b int) int { return a / b },
}

type monkey struct {
	value *int
	op    *operation
}

func (m monkey) String() string {
	if m.value != nil {
		return fmt.Sprintf("%d", *m.value)
	} else {
		return fmt.Sprintf("%s %c %s", m.op.left, m.op.operator, m.op.right)
	}
}

type day21 struct {
	monkeys map[string]monkey
}

func (d *day21) evaluate(name string) (int, error) {
	monkey, ok := d.monkeys[name]
	if !ok {
		return 0, fmt.Errorf("monkey %s not found", name)
	}
	if monkey.value != nil {
		return *monkey.value, nil
	}

	op := monkey.op
	left, err := d.evaluate(op.left)
	if err != nil {
		return 0, err
	}
	right, err := d.evaluate(op.right)
	if err != nil {
		return 0, err
	}
	return operators[op.operator](left, right), nil
}

func (d *day21) solveForHumn(name string, wanted int) int {
	monkey, ok := d.monkeys[name]
	if !ok {
		return wanted
	}

	left, leftErr := d.evaluate(monkey.op.left)
	right, rightErr := d.evaluate(monkey.op.right)

	if leftErr != nil {
		// humn is in the left
		switch monkey.op.operator {
		case '+':
			return d.solveForHumn(monkey.op.left, wanted-right)
		case '-':
			return d.solveForHumn(monkey.op.left, wanted+right)
		case '*':
			return d.solveForHumn(monkey.op.left, wanted/right)
		case '/':
			return d.solveForHumn(monkey.op.left, wanted*right)
		}
	}

	if rightErr != nil {
		// humn is in the right
		switch monkey.op.operator {
		case '+':
			return d.solveForHumn(monkey.op.right, wanted-left)
		case '-':
			return d.solveForHumn(monkey.op.right, left-wanted)
		case '*':
			return d.solveForHumn(monkey.op.right, wanted/left)
		case '/':
			return d.solveForHumn(monkey.op.right, left*wanted)
		}
	}

	return -1
}

func (d *day21) Part1() int {
	return utils.Must(d.evaluate("root"))
}

func (d *day21) Part2() int {
	delete(d.monkeys, "humn")

	var (
		rootRightExpression, rightErr = d.evaluate(d.monkeys["root"].op.right)
		rootLeftExpression, leftErr   = d.evaluate(d.monkeys["root"].op.left)
	)

	if rightErr != nil {
		// This one has humn in it
		return d.solveForHumn(d.monkeys["root"].op.right, rootLeftExpression)
	} else if leftErr != nil {
		// This one has humn in it
		return d.solveForHumn(d.monkeys["root"].op.left, rootRightExpression)
	}

	return 0
}

func Parse(filename string) *day21 {
	data := utils.ReadLines(filename)

	monkeys := make(map[string]monkey, len(data))

	for _, line := range data {
		splitLine := strings.Split(line, ": ")
		right := strings.Split(splitLine[1], " ")
		if len(right) == 1 {
			// This is a value
			val := utils.Must(strconv.Atoi(right[0]))
			monkeys[splitLine[0]] = monkey{value: &val}
		} else {
			// This is an operation
			op := operation{
				left:     right[0],
				right:    right[2],
				operator: rune(right[1][0]),
			}
			monkeys[splitLine[0]] = monkey{op: &op}
		}
	}

	return &day21{monkeys: monkeys}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: root's number:", day.Part1())
	fmt.Println("ANSWER2: number that human should say:", day.Part2())
}
