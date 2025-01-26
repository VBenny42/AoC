package day08

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type direction int

const (
	left direction = iota
	right
)

type node struct {
	name        string
	left, right *node
}

type day08 struct {
	directions []direction
	nodes      map[string]*node
	startNodes []*node
}

func findNode(nodes map[string]*node, name string) *node {
	if n, ok := nodes[name]; ok {
		return n
	}

	n := &node{name: name}
	nodes[name] = n
	return n
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (gcd) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (lcm) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := range integers {
		result = lcm(result, integers[i])
	}

	return result
}

func (d *day08) Part1() (steps int) {
	var (
		start   = d.nodes["AAA"]
		end     = d.nodes["ZZZ"]
		current = start
		dirIdx  int
	)

	for {
		dir := d.directions[dirIdx]
		dirIdx = (dirIdx + 1) % len(d.directions)

		if dir == left {
			current = current.left
		} else {
			current = current.right
		}

		steps++

		if current == end {
			break
		}
	}

	return
}

func (d *day08) Part2() int {
	startToEndLengths := make([]int, len(d.startNodes))

	for i, n := range d.startNodes {
		var (
			current = n
			dirIdx  int
			steps   int
		)

		for {
			dir := d.directions[dirIdx]
			dirIdx = (dirIdx + 1) % len(d.directions)

			switch dir {
			case left:
				current = current.left
			case right:
				current = current.right
			}

			steps++

			if strings.HasSuffix(current.name, "Z") {
				startToEndLengths[i] = steps
				break
			}
		}
	}

	return lcm(startToEndLengths[0], startToEndLengths[1], startToEndLengths[2:]...)
}

func Parse(filename string) *day08 {
	data := utils.ReadLines(filename)

	var (
		nodes      = make(map[string]*node, len(data[2:]))
		directions = make([]direction, len(data[0]))
		startNodes []*node
	)

	for i, c := range data[0] {
		switch c {
		case 'L':
			directions[i] = left
		case 'R':
			directions[i] = right
		}
	}

	var node, leftNode, rightNode *node

	for _, line := range data[2:] {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ',' || r == '(' || r == ')'
		})
		if len(fields) != 4 {
			fmt.Println("Error parsing line:", line)
			fmt.Println("fields:", fields)
			panic("Invalid line")
		}

		node = findNode(nodes, fields[0])
		leftNode = findNode(nodes, fields[2])
		rightNode = findNode(nodes, fields[3])

		node.left = leftNode
		node.right = rightNode

		if strings.HasSuffix(fields[0], "A") {
			startNodes = append(startNodes, node)
		}
	}

	return &day08{directions, nodes, startNodes}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: steps to reach ZZZ from AAA:", day.Part1())
	fmt.Println("ANSWER2: steps to land on all Z nodes:", day.Part2())
}
