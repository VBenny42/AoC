package utils

import (
	"strings"

	"github.com/VBenny42/AoC/2024/golang/embeds"
)

func SplitLines(filename string) []string {
	file, err := embeds.Inputs.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.TrimSpace(string(file))

	return strings.Split(lines, "\n")
}

func JoinFile(filename string) string {
	file, err := embeds.Inputs.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(file))
}
