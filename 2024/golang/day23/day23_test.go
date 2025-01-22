package day23_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day23"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day23.Parse("inputs/day23/sample-input.txt")

	value, _ := day.Part1and2()

	assert.Equal(t, 7, value)
}

func TestPart2(t *testing.T) {
	day := day23.Parse("inputs/day23/sample-input.txt")

	_, value := day.Part1and2()

	assert.Equal(t, "co,de,ka,ta", value)
}
