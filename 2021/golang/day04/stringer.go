package day04

import (
	"fmt"
	"strings"
)

func (b board) String() string {
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			builder.WriteString(fmt.Sprintf("% 2d ", b.numbers[i][j]))
		}
		builder.WriteString("\n")
	}
	builder.WriteString("\n")

	return builder.String() + fmt.Sprintf("Occupied: %v\n", b.occupied)
}
