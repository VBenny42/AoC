package day12_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day12"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day12.Parse("inputs/day12/sample-input.txt")

	assert.Equal(t, 21, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day12.Parse("inputs/day12/sample-input.txt")

	assert.Equal(t, 525152, d.Part2())
}
