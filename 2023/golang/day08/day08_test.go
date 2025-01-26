package day08_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day08"
	"github.com/stretchr/testify/assert"
)

func TestPart1RL(t *testing.T) {
	d := day08.Parse("inputs/day08/sample-input.txt")

	assert.Equal(t, 2, d.Part1())
}

func TestPart1LLR(t *testing.T) {
	d := day08.Parse("inputs/day08/sample-input-1.txt")

	assert.Equal(t, 6, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day08.Parse("inputs/day08/sample-input-2.txt")

	assert.Equal(t, 6, d.Part2())
}
