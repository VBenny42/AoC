package day05_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day05"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	d := day05.Parse("../../inputs/day05/sample-input.txt")

	assert.Equal(t, 3, len(d.Stacks))

	assert.Equal(t, 4, len(d.Instructions))

	assert.Equal(t, []rune{'Z', 'N'}, d.Stacks[0])
	assert.Equal(t, []rune{'M', 'C', 'D'}, d.Stacks[1])
	assert.Equal(t, []rune{'P'}, d.Stacks[2])
}

func TestPart1(t *testing.T) {
	d := day05.Parse("../../inputs/day05/sample-input.txt")

	assert.Equal(t, "CMZ", d.Part1())
}

func TestPart2(t *testing.T) {
	d := day05.Parse("../../inputs/day05/sample-input.txt")

	assert.Equal(t, "MCD", d.Part2())
}
