//go:build ignore

package day24

import "fmt"

func (p point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.x, p.y, p.z)
}

func (h hailstone) String() string {
	return fmt.Sprintf("%s @ %s", h.position, h.velocity)
}
