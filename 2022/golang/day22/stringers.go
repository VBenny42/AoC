package day22

import (
	"fmt"
	"strings"
)

func (g grid) String() string {
	var builder strings.Builder
	for _, row := range g {
		for _, cell := range row {
			switch {
			case cell&(1<<last) != 0:
				builder.WriteRune('X')
			case cell&(1<<wall) != 0:
				builder.WriteRune('#')
			case cell&(1<<up) != 0:
				builder.WriteRune('^')
			case cell&(1<<down) != 0:
				builder.WriteRune('v')
			case cell&(1<<left) != 0:
				builder.WriteRune('<')
			case cell&(1<<right) != 0:
				builder.WriteRune('>')
			case cell&(1<<movable) != 0:
				builder.WriteRune('.')
			default:
				builder.WriteRune(' ')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func (m movement) String() string {
	if m.amount != nil {
		return fmt.Sprintf("Move %d steps", *m.amount)
	}
	return fmt.Sprintf("Turn %c", *m.rotation)
}
