package day20_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day20"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day20.Parse("inputs/day20/sample-input.txt")

	assert.Equal(t, 3, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day20.Parse("inputs/day20/sample-input.txt")

	assert.Equal(t, 1623178306, day.Part2())
}
