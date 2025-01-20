package day01_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day01"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day01.Parse("inputs/day01/sample-input.txt")

	assert.Equal(t, 24000, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day01.Parse("inputs/day01/sample-input.txt")

	assert.Equal(t, 45000, d.Part2())
}
