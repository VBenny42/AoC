package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type op rune

const (
	plus     op = '+'
	multiply op = '*'
)

type operation struct {
	operator op
	operand  int
	selfOp   bool
}

type monkey struct {
	num                   int
	items                 []int
	operation             operation
	test                  int
	trueThrow, falseThrow int
	times                 int
}

type day11 struct {
	monkeys []monkey
	modBy   int
}

func (d *day11) round(divideAmount int) {
	shouldDivide := divideAmount == 1
	for i := range d.monkeys {
		d.monkeys[i].times += len(d.monkeys[i].items)
		for _, item := range d.monkeys[i].items {
			worryLevel := d.monkeys[i].operation.operand
			if d.monkeys[i].operation.selfOp {
				worryLevel = item
			}
			switch d.monkeys[i].operation.operator {
			case plus:
				worryLevel += item
			case multiply:
				worryLevel *= item
			}
			worryLevel /= divideAmount

			if shouldDivide {
				worryLevel %= d.modBy
			}

			if worryLevel%d.monkeys[i].test == 0 {
				d.monkeys[d.monkeys[i].trueThrow].items = append(
					d.monkeys[d.monkeys[i].trueThrow].items, worryLevel)
			} else {
				d.monkeys[d.monkeys[i].falseThrow].items = append(
					d.monkeys[d.monkeys[i].falseThrow].items, worryLevel)
			}
		}
		d.monkeys[i].items = d.monkeys[i].items[:0]
	}
}

func (d *day11) calcMonkeyBusiness(rounds, divideAmount int) int {
	for range rounds {
		d.round(divideAmount)
	}

	sort.Slice(d.monkeys, func(i, j int) bool {
		return d.monkeys[i].times > d.monkeys[j].times
	})

	monkeyBusiness := d.monkeys[0].times * d.monkeys[1].times

	return monkeyBusiness
}

func (d *day11) Part1() int {
	return d.calcMonkeyBusiness(20, 3)
}

func (d *day11) Part2() int {
	return d.calcMonkeyBusiness(10000, 1)
}

func Parse(filename string) *day11 {
	data := utils.ReadLines(filename)
	monkeyCount := (len(data) + 1) / 7
	monkeys := make([]monkey, monkeyCount)
	modBy := 1

	for i := 0; i < len(data); i += 7 {
		if i >= len(data) {
			break
		}

		num := i / 7
		monkeys[num].num = num

		// Parse starting items
		itemsStr := strings.Split(strings.TrimSpace(strings.Split(data[i+1], ":")[1]), ", ")
		monkeys[num].items = make([]int, len(itemsStr))
		for j, item := range itemsStr {
			monkeys[num].items[j] = utils.Must(strconv.Atoi(item))
		}

		// Parse operation
		opParts := strings.Split(strings.TrimSpace(strings.Split(data[i+2], "=")[1]), " ")
		monkeys[num].operation.operator = op(opParts[1][0])
		operand, err := strconv.Atoi(opParts[2])
		if err != nil {
			monkeys[num].operation.selfOp = true
		}
		monkeys[num].operation.operand = operand

		// Parse test
		testParts := strings.Split(data[i+3], "divisible by ")
		monkeys[num].test = utils.Must(strconv.Atoi(strings.TrimSpace(testParts[1])))
		modBy *= monkeys[num].test

		// Parse true/false cases
		monkeys[num].trueThrow = utils.Must(strconv.Atoi(
			strings.TrimSpace(strings.Split(data[i+4], "monkey ")[1])))
		monkeys[num].falseThrow = utils.Must(strconv.Atoi(
			strings.TrimSpace(strings.Split(data[i+5], "monkey ")[1])))
	}

	return &day11{monkeys, modBy}
}

func Solve(filename string) {
	fmt.Println("ANSWER1: monkey business after 20 rounds", Parse(filename).Part1())
	fmt.Println("ANSWER2: monkey business after 10000 rounds", Parse(filename).Part2())
}
