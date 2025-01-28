package day12

import (
	"strconv"
	"strings"
)

func (r row) String() string {
	var builder strings.Builder

	for _, c := range r.springs {
		switch c {
		case operational:
			builder.WriteRune('.')
		case damaged:
			builder.WriteRune('#')
		case unknown:
			builder.WriteRune('?')
		default:
			panic("invalid condition")
		}
	}

	builder.WriteRune(' ')

	for i, d := range r.groups {
		if i > 0 && i < len(r.groups) {
			builder.WriteRune(',')
		}
		builder.WriteString(strconv.Itoa(d))
	}

	return builder.String()
}
