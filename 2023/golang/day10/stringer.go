package day10

import "strings"

func (g grid) String() string {
	var builder strings.Builder

	for _, row := range g.grid {
		builder.WriteString(string(row) + "\n")
	}

	return builder.String()
}
