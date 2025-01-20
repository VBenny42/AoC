package day24

import "strings"

func (g grid) String() string {
	var builder strings.Builder
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			switch g[y][x] {
			case true:
				builder.WriteRune('#')
			case false:
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}
