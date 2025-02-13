package day18

import (
	"fmt"
	"strconv"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

// Gave up on this one, solution from reddit
type day18 struct {
	pairs []string
}

type number struct {
	value       int
	left, right *number
	parent      *number
	isRegular   bool
}

func newRegular(value int) *number {
	return &number{value: value, isRegular: true}
}

func newPair(left, right *number) *number {
	node := &number{left: left, right: right, isRegular: false}
	left.parent = node
	right.parent = node
	return node
}

func parseSnailfish(s string) *number {
	if len(s) == 0 {
		return nil
	}

	// Check if it's a regular number
	if n, err := strconv.Atoi(s); err == nil {
		return newRegular(n)
	}

	// Remove outer brackets
	s = s[1 : len(s)-1]

	// Find the middle comma at the current level
	depth := 0
	for i, c := range s {
		if c == '[' {
			depth++
		} else if c == ']' {
			depth--
		} else if c == ',' && depth == 0 {
			left := parseSnailfish(s[:i])
			right := parseSnailfish(s[i+1:])
			return newPair(left, right)
		}
	}
	return nil
}

func add(a, b *number) *number {
	if a == nil {
		return b
	}
	return newPair(a, b)
}

func (n *number) magnitude() int {
	if n.isRegular {
		return n.value
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func findSplit(node *number) *number {
	if node == nil {
		return nil
	}
	if node.isRegular && node.value >= 10 {
		return node
	}
	if !node.isRegular {
		if left := findSplit(node.left); left != nil {
			return left
		}
		return findSplit(node.right)
	}
	return nil
}

func (n *number) split() {
	if !n.isRegular {
		return
	}
	leftVal := n.value / 2
	rightVal := (n.value + 1) / 2
	n.left = newRegular(leftVal)
	n.right = newRegular(rightVal)
	n.left.parent = n
	n.right.parent = n
	n.isRegular = false
	n.value = 0
}

func findExplode(node *number, depth int) *number {
	if node == nil {
		return nil
	}
	if depth >= 4 && !node.isRegular && node.left.isRegular && node.right.isRegular {
		return node
	}
	if !node.isRegular {
		if left := findExplode(node.left, depth+1); left != nil {
			return left
		}
		return findExplode(node.right, depth+1)
	}
	return nil
}

func (n *number) findNextLeft() *number {
	current := n
	parent := current.parent
	for parent != nil && parent.left == current {
		current = parent
		parent = current.parent
	}
	if parent == nil {
		return nil
	}
	current = parent.left
	for !current.isRegular {
		current = current.right
	}
	return current
}

func (n *number) findNextRight() *number {
	current := n
	parent := current.parent
	for parent != nil && parent.right == current {
		current = parent
		parent = current.parent
	}
	if parent == nil {
		return nil
	}
	current = parent.right
	for !current.isRegular {
		current = current.left
	}
	return current
}

func (n *number) explode() {
	if n.isRegular {
		return
	}
	leftVal := n.left.value
	rightVal := n.right.value

	if left := n.findNextLeft(); left != nil {
		left.value += leftVal
	}
	if right := n.findNextRight(); right != nil {
		right.value += rightVal
	}

	n.isRegular = true
	n.value = 0
	n.left = nil
	n.right = nil
}

func reduce(node *number) bool {
	if explodeNode := findExplode(node, 0); explodeNode != nil {
		explodeNode.explode()
		return true
	}
	if splitNode := findSplit(node); splitNode != nil {
		splitNode.split()
		return true
	}
	return false
}

func (d *day18) Part1() int {
	var root *number
	for _, pair := range d.pairs {
		number := parseSnailfish(pair)
		root = add(root, number)
		for reduce(root) {
		}
	}
	return root.magnitude()
}

func (d *day18) Part2() int {
	var (
		maxMagnitude int
		sum          *number
	)

	addPair := func(i, j int) {
		sum = add(parseSnailfish(d.pairs[i]), parseSnailfish(d.pairs[j]))
		for reduce(sum) {
		}
		maxMagnitude = max(maxMagnitude, sum.magnitude())
	}

	for i := 0; i < len(d.pairs); i++ {
		for j := 0; j < len(d.pairs); j++ {
			if i == j {
				continue
			}

			addPair(i, j)
			addPair(j, i)
		}
	}
	return maxMagnitude
}

func Parse(filename string) *day18 {
	return &day18{pairs: utils.ReadLines(filename)}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: magnitude of final sum:", day.Part1())
	fmt.Println("ANSWER2: maximum magnitude of any two different pairs:", day.Part2())
}
