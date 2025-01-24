package utils

import (
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/embeds"
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

func Atoi(s string) int {
	return Must(strconv.Atoi(s))
}
