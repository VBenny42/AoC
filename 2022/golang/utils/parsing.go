package utils

import (
	"strings"

	"github.com/VBenny42/AoC/2022/golang/embeds"
)

func ReadLines(filename string) []string {
	file := Must(embeds.Inputs.ReadFile(filename))

	lines := strings.Trim(string(file), "\n")

	return strings.Split(lines, "\n")
}

func ReadTrimmed(filename string) string {
	file := Must(embeds.Inputs.ReadFile(filename))

	return strings.Trim(string(file), "\n")
}
