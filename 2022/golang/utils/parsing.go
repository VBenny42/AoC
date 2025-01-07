package utils

import (
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	file := Must(os.ReadFile(filename))

	lines := strings.Trim(string(file), "\n")

	return strings.Split(lines, "\n")
}

func ReadTrimmed(filename string) string {
	file := Must(os.ReadFile(filename))

	return strings.Trim(string(file), "\n")
}
