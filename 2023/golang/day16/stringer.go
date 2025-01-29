package day16

import "strings"

func (g grid) String() string {
	var builder strings.Builder

	for _, row := range g {
		for _, cell := range row {
			switch {
			case cell&(1<<vSplit) != 0:
				builder.WriteRune('|')
			case cell&(1<<hSplit) != 0:
				builder.WriteRune('-')
			case cell&(1<<rightMirror) != 0:
				builder.WriteRune('/')
			case cell&(1<<leftMirror) != 0:
				builder.WriteRune('\\')
			case cell == 0:
				builder.WriteRune('.')
			default:
				builder.WriteRune('?')
			}
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}
