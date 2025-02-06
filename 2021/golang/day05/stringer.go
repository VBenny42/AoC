package day05

import (
	"strconv"
	"strings"
)

func printGrid(grid grid) string {
	var builder strings.Builder

	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[y][x] == 0 {
				builder.WriteString(".")
				continue
			}
			builder.WriteString(strconv.Itoa(grid[y][x]))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
