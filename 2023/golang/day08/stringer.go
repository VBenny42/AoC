package day08

import "fmt"

func (d direction) String() string {
	switch d {
	case left:
		return "L"
	case right:
		return "R"
	}
	panic("unknown direction")
}

func (n *node) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%s = (%s, %s)", n.name, n.left.name, n.right.name)
}
