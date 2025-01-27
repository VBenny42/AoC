package day11

import (
	"strings"
)

func (g grid) String() string {
	var builder strings.Builder

	for _, row := range g {
		for _, cell := range row {
			switch cell {
			case galaxy:
				builder.WriteRune('#')
			case empty:
				builder.WriteRune('.')
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
