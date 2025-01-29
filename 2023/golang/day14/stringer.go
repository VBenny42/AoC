package day14

import "strings"

func (g grid) String() string {
	var builder strings.Builder

	for _, row := range g {
		for _, cell := range row {
			builder.WriteRune(rune(cell))
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}
