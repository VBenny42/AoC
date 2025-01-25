package day03_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day03"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day03.Parse("inputs/day03/sample-input.txt")

	assert.Equal(t, 4361, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day03.Parse("inputs/day03/sample-input.txt")

	assert.Equal(t, 467835, d.Part2())
}
