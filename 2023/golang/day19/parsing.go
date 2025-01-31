package day19

import (
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

func parseCondition(str string) *condition {
	var (
		idx   partIdx
		op    operator
		value int
	)

	switch str[0] {
	case 'x':
		idx = x
	case 'm':
		idx = m
	case 'a':
		idx = a
	case 's':
		idx = s
	}

	switch str[1] {
	case '>':
		op = gt
	case '<':
		op = lt
	}

	value = utils.Atoi(str[2:])

	return &condition{idx, op, value}
}

func Parse(filename string) *day19 {
	var (
		data      = strings.Split(utils.ReadTrimmed(filename), "\n\n")
		workflows = make(map[string]workflow, len(data[0])+2)
		partsStr  = strings.Split(data[1], "\n")
		parts     = make([]part, len(partsStr))
	)

	// Special accepted case
	workflows["A"] = workflow{}
	// Special rejected case
	workflows["R"] = workflow{}

	splitWorkflow := func(r rune) bool {
		return r == ',' || r == '{' || r == '}'
	}

	for _, line := range strings.Split(data[0], "\n") {
		var (
			fields   = strings.FieldsFunc(line, splitWorkflow)
			name     = fields[0]
			rulesStr = fields[1:]
			wf       = workflow{
				rules: make([]rule, len(rulesStr)),
			}
		)

		workflows[name] = wf

		for i, ruleStr := range rulesStr {
			parts := strings.SplitN(ruleStr, ":", 2)
			wf.rules[i].target = parts[0]
			if len(parts) == 2 {
				wf.rules[i].cond = parseCondition(parts[0])
				wf.rules[i].target = parts[1]
			}
		}
	}

	for i, line := range partsStr {
		line = strings.Trim(line, "{}")

		var p part

		for _, comp := range strings.Split(line, ",") {
			var (
				pair  = strings.SplitN(comp, "=", 2)
				key   = pair[0]
				value = utils.Atoi(pair[1])
			)
			switch key {
			case "x":
				p[x] = value
			case "m":
				p[m] = value
			case "a":
				p[a] = value
			case "s":
				p[s] = value
			}
		}

		parts[i] = p
	}

	return &day19{
		workflows: workflows,
		parts:     parts,
	}
}
