package day18

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

func (i instruction) String() string {
	var s string

	switch i.direction {
	case utils.Up:
		s = "U"
	case utils.Down:
		s = "D"
	case utils.Left:
		s = "L"
	case utils.Right:
		s = "R"
	}

	s += " "

	s += strconv.Itoa(i.distance)

	s += " "

	s += fmt.Sprintf("#%02x%02x%02x", i.color.R, i.color.G, i.color.B)

	return s
}

func (g grid) String() string {
	var builder strings.Builder

	var sentinel color.RGBA

	for _, row := range g {
		for _, cell := range row {
			if cell != sentinel {
				builder.WriteRune('█')
			} else {
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}
