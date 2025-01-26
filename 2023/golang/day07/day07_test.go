package day07_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day07"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day07.Parse("inputs/day07/sample-input.txt")

	assert.Equal(t, 6440, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day07.Parse("inputs/day07/sample-input.txt")

	assert.Equal(t, 5905, d.Part2())
}

func TestPart1Harder(t *testing.T) {
	d := day07.Parse("inputs/day07/sample-input-harder.txt")

	assert.Equal(t, 6592, d.Part1())
}

func TestPart2Harder(t *testing.T) {
	d := day07.Parse("inputs/day07/sample-input-harder.txt")

	assert.Equal(t, 6839, d.Part2())
}
