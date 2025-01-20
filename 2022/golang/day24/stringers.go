package day24

import "strings"

func (g grid) String() string {
	var builder strings.Builder
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			builder.WriteRune(g[y][x])
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}
