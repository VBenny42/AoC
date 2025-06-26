package day16

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
	"github.com/VBenny42/AoC/2020/golang/utils/set"
)

type (
	ticket     []int
	validRange struct {
		start, end int
	}
	rule struct {
		v1, v2 validRange
	}
)

type day16 struct {
	rules         map[string]rule
	myTicket      ticket
	nearbyTickets []ticket
	validTickets  []ticket
}

func (d *day16) Part1() (errorRate int) {
	invalidTickets := set.NewSet[int]()

	for i, t := range d.nearbyTickets {
		for _, v := range t {
			valid := false
			for _, r := range d.rules {
				if (v >= r.v1.start && v <= r.v1.end) || (v >= r.v2.start && v <= r.v2.end) {
					valid = true
					break
				}
			}
			if !valid {
				errorRate += v
				invalidTickets.Add(i)
			}
		}
	}

	for i, t := range d.nearbyTickets {
		if !invalidTickets.Contains(i) {
			d.validTickets = append(d.validTickets, t)
		}
	}

	return
}

func (d *day16) Part2() int {
	length := len(d.myTicket)

	possibleColumns := make(map[string][]int)

	for ruleName, rule := range d.rules {
		for i := 0; i < length; i++ {
			valid := true
			for _, t := range d.validTickets {
				if (t[i] < rule.v1.start || t[i] > rule.v1.end) && (t[i] < rule.v2.start || t[i] > rule.v2.end) {
					valid = false
					break
				}
			}
			if valid {
				possibleColumns[ruleName] = append(possibleColumns[ruleName], i)
			}
		}
	}

	fieldToColumn := make(map[string]int)
	usedColumns := set.NewSet[int]()

	changed := true
	for changed {
		changed = false

		for ruleName, columns := range possibleColumns {
			if len(columns) == 1 {
				col := columns[0]
				if !usedColumns.Contains(col) {
					fieldToColumn[ruleName] = col
					usedColumns.Add(col)
					changed = true

					for otherRule, otherColumns := range possibleColumns {
						if otherRule != ruleName {
							newColumns := []int{}
							for _, c := range otherColumns {
								if c != col {
									newColumns = append(newColumns, c)
								}
							}
							possibleColumns[otherRule] = newColumns
						}
					}
				}
			}
		}
	}

	result := 1
	for fieldName, columnIndex := range fieldToColumn {
		if strings.HasPrefix(fieldName, "departure") {
			result *= d.myTicket[columnIndex]
		}
	}

	return result
}

func Parse(filename string) *day16 {
	lines := utils.ReadLines(filename)
	rules := make(map[string]rule)

	var myTicket ticket
	var nearbyTickets []ticket

	var (
		rulesDone         bool
		myTicketDone      bool
		nearbyTicketsLine int
	)
	var r1, r2 validRange

	for i := 0; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" {
			if !rulesDone {
				rulesDone = true
			} else if !myTicketDone {
				myTicketDone = true
			}
			continue
		}

		if !rulesDone {
			left, right, ok := strings.Cut(lines[i], ": ")
			if !ok {
				panic(fmt.Sprintf("could not parse rule line %d: %s", i, lines[i]))
			}
			split := strings.Split(right, " or ")
			if len(split) != 2 {
				panic(fmt.Sprintf("could not parse rule line %d: %s", i, lines[i]))
			}
			split[0] = strings.TrimSpace(split[0])
			n, err := fmt.Sscanf(split[0], "%d-%d", &r1.start, &r1.end)
			if err != nil || n != 2 {
				panic(fmt.Sprintf("could not parse range in rule line %d: %s", i, split[0]))
			}
			split[1] = strings.TrimSpace(split[1])
			n, err = fmt.Sscanf(split[1], "%d-%d", &r2.start, &r2.end)
			if err != nil || n != 2 {
				panic(fmt.Sprintf("could not parse range in rule line %d: %s", i, split[1]))
			}

			rules[left] = rule{v1: r1, v2: r2}
		}

		if (!myTicketDone) && (i > 0) && (strings.HasPrefix(lines[i-1], "your ticket:")) {
			split := strings.Split(lines[i], ",")
			myTicket = make(ticket, len(split))
			for j, s := range split {
				myTicket[j] = utils.Atoi(s)
			}
			myTicketDone = true
			continue
		}

		if strings.HasPrefix(lines[i], "nearby tickets:") {
			nearbyTicketsLine = i + 1
		}
		if rulesDone && myTicketDone && (i >= nearbyTicketsLine) {
			split := strings.Split(lines[i], ",")
			t := make(ticket, len(split))
			for j, s := range split {
				t[j] = utils.Atoi(s)
			}
			nearbyTickets = append(nearbyTickets, t)
		}
	}
	return &day16{rules: rules, myTicket: myTicket, nearbyTickets: nearbyTickets}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: ticket scanning error rate:", day.Part1())
	fmt.Println("ANSWER2: result of departure fields:", day.Part2())
}
