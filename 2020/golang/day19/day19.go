package day19

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day19 struct {
	rules    map[string]string
	messages []string
}

func (d *day19) buildRule(rule string) string {
	if strings.HasPrefix(rule, "\"") && strings.HasSuffix(rule, "\"") {
		return rule[1 : len(rule)-1]
	}
	if strings.Contains(rule, "|") {
		parts := strings.Split(rule, " | ")
		var options []string
		for _, part := range parts {
			subRules := strings.Split(part, " ")
			var subRuleParts []string
			for _, subRule := range subRules {
				subRuleParts = append(subRuleParts, d.buildRule(d.rules[subRule]))
			}
			options = append(options, strings.Join(subRuleParts, ""))
		}
		return "(" + strings.Join(options, "|") + ")"
	} else {
		subRules := strings.Split(rule, " ")
		var subRuleParts []string
		for _, subRule := range subRules {
			subRuleParts = append(subRuleParts, d.buildRule(d.rules[subRule]))
		}
		return strings.Join(subRuleParts, "")
	}
}

func (d *day19) Part1() (count int) {
	rule0 := d.buildRule(d.rules["0"])

	pattern := "^" + rule0 + "$"
	regex := regexp.MustCompile(pattern)

	for _, message := range d.messages {
		if regex.MatchString(message) {
			count++
		}
	}

	return
}

// Made with Claude hehe
func (d *day19) Part2() (count int) {
	// Create a copy of rules to avoid modifying original
	rulesCopy := make(map[string]string)
	for k, v := range d.rules {
		rulesCopy[k] = v
	}

	// Build rule 42 and 31 first
	rule42 := d.buildRule(d.rules["42"])
	rule31 := d.buildRule(d.rules["31"])

	// Rule 8: 42 | 42 8 becomes 42+ (one or more 42s)
	// Rule 11: 42 31 | 42 11 31 becomes balanced pairs
	// Since regex can't handle true balanced pairs, we approximate with limited depth

	var rule11Options []string
	// Generate patterns for rule 11: 42{n} 31{n} for n = 1 to some reasonable limit
	for i := 1; i <= 10; i++ { // Adjust limit as needed
		pattern := strings.Repeat("("+rule42+")", i) + strings.Repeat("("+rule31+")", i)
		rule11Options = append(rule11Options, pattern)
	}
	rule11Pattern := "(" + strings.Join(rule11Options, "|") + ")"

	// Rule 0: 8 11 becomes 42+ followed by balanced 42{n}31{n}
	rule0Pattern := "^(" + rule42 + ")+" + rule11Pattern + "$"

	regex := regexp.MustCompile(rule0Pattern)
	for _, message := range d.messages {
		if regex.MatchString(message) {
			count++
		}
	}
	return
}

func Parse(filename string) *day19 {
	lines := utils.ReadLines(filename)
	rules := make(map[string]string)
	var messages []string

	rulesDone := false

	for _, line := range lines {
		if line == "" {
			rulesDone = true
			continue
		}

		if !rulesDone {
			parts := strings.Split(line, ": ")
			rules[parts[0]] = parts[1]
		} else {
			messages = append(messages, line)
		}
	}

	return &day19{rules: rules, messages: messages}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of messages that completely match rule 0:", day.Part1())
	fmt.Println("ANSWER2: number of messages that completely match rule 0 after modifying rules 8 and 11:", day.Part2())
}
