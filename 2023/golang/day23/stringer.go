package day23

import "strings"

func (g grid) String() string {
	cellToRune := map[cell]rune{
		empty:      '.',
		leftSlope:  '<',
		rightSlope: '>',
		upSlope:    '^',
		downSlope:  'v',
		wall:       '#',
	}

	var builder strings.Builder
	for _, row := range g {
		for _, c := range row {
			val, ok := cellToRune[c]
			if !ok {
				builder.WriteRune('?')
				continue
			}
			builder.WriteRune(val)
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}
