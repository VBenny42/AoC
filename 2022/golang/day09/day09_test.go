package day09_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day09"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	day := day09.Parse("../../inputs/day09/sample-input.txt")

	check := []day09.Motion{
		{Direction: 'R', Distance: 4},
		{Direction: 'U', Distance: 4},
		{Direction: 'L', Distance: 3},
		{Direction: 'D', Distance: 1},
		{Direction: 'R', Distance: 4},
		{Direction: 'D', Distance: 1},
		{Direction: 'L', Distance: 5},
		{Direction: 'R', Distance: 2},
	}

	assert.Equal(t, check, day.Motions)
}

func TestPart1(t *testing.T) {
	day := day09.Parse("../../inputs/day09/sample-input.txt")

	assert.Equal(t, 13, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day09.Parse("../../inputs/day09/sample-input-larger.txt")

	assert.Equal(t, 36, day.Part2())
}
