package utils

import (
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/embeds"
)

func ReadLines(filename string) []string {
	file := Must(embeds.Inputs.ReadFile(filename))

	lines := strings.Trim(string(file), "\n")

	return strings.Split(lines, "\n")
}

// func ReadLines(filename string) iter.Seq[string] {
// 	file := Must(embeds.Inputs.ReadFile(filename))
//
// 	return strings.Lines(string(file))
// }

func ReadTrimmed(filename string) string {
	file := Must(embeds.Inputs.ReadFile(filename))

	return strings.Trim(string(file), "\n")
}

func Atoi(s string) int {
	return Must(strconv.Atoi(s))
}
