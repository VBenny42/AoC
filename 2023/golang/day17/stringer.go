package day17

import (
	"strconv"
	"strings"
)

func (g grid) String() string {
	var builder strings.Builder

	for _, row := range g {
		for _, cell := range row {
			builder.WriteString(strconv.Itoa(cell))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
