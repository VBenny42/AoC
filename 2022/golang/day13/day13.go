package day13

import (
	"fmt"
	// "sort"
	"strconv"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type packet struct {
	value    *int
	children []packet
}

type pair struct {
	left, right packet
}

type day13 struct {
	pairs []pair
}

func (packet packet) String() string {
	if packet.value != nil {
		return strconv.Itoa(*packet.value)
	}

	result := "["
	for i, child := range packet.children {
		if i > 0 {
			result += ","
		}
		result += child.String()
	}
	result += "]"
	return result
}

func (pair pair) String() string {
	return fmt.Sprintf("{%s, %s}", pair.left, pair.right)
}

func parsePacket(s string) packet {
	if s == "[]" {
		return packet{children: []packet{}}
	}

	if s[0] != '[' {
		val := utils.Must(strconv.Atoi(s))
		return packet{value: &val}
	}

	// Remove outer brackets
	s = s[1 : len(s)-1]

	var result packet
	result.children = make([]packet, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			// Find matching closing bracket
			depth := 1
			j := i + 1
			for depth > 0 {
				if s[j] == '[' {
					depth++
				} else if s[j] == ']' {
					depth--
				}
				j++
			}
			result.children = append(result.children, parsePacket(s[i:j]))
			i = j
			if i < len(s) && s[i] == ',' {
				i++
			}
			i--
			continue
		}

		if s[i] == ',' {
			continue
		}

		// Parse number
		j := i
		for j < len(s) && s[j] >= '0' && s[j] <= '9' {
			j++
		}
		val := utils.Must(strconv.Atoi(s[i:j]))
		result.children = append(result.children, packet{value: &val})
		i = j - 1
	}

	return result
}

func comparePair(left, right packet) (bool, error) {
	minLen := len(left.children)
	if len(right.children) < minLen {
		minLen = len(right.children)
	}

	for i := 0; i < minLen; i++ {
		l, r := left.children[i], right.children[i]

		if l.value != nil && r.value != nil {
			if *l.value > *r.value {
				return false, nil
			}
			if *l.value < *r.value {
				return true, nil
			}
			continue
		}

		if l.value != nil {
			l = packet{children: []packet{{value: l.value}}}
		}
		if r.value != nil {
			r = packet{children: []packet{{value: r.value}}}
		}

		if result, err := comparePair(l, r); err == nil {
			return result, nil
		}
	}

	if len(left.children) < len(right.children) {
		return true, nil
	}
	if len(left.children) > len(right.children) {
		return false, nil
	}

	return false, fmt.Errorf("undecided")
}

func (d *day13) Part1() int {
	sum := 0

	for i, pair := range d.pairs {
		if result, err := comparePair(pair.left, pair.right); err == nil && result {
			sum += i + 1
		}
	}

	return sum
}

// func (d *day13) Part2() int {
// 	d.pairs = append(d.pairs, )
// 	sort.Slice(d.pairs, func(i, j int) bool {
// 		result, err := comparePair(d.pairs[i].left, d.pairs[j].left)
// 		if err != nil {
// 			return false
// 		}
// 		return result
// 	})
//
// 	return 0
// }

func Parse(filename string) *day13 {
	data := utils.ReadLines(filename)

	pairs := make([]pair, 0)

	var pair pair

	for i := 0; i < len(data); i += 3 {
		if i >= len(data) {
			break
		}

		pair.left = parsePacket(data[i])
		pair.right = parsePacket(data[i+1])

		pairs = append(pairs, pair)
	}

	return &day13{pairs}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: pairs in the right order:", day.Part1())
}
