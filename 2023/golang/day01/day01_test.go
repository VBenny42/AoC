package day01_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day01"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day01.Parse("inputs/day01/sample-input.txt")

	assert.Equal(t, 142, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day01.Parse("inputs/day01/sample-input-2.txt")

	assert.Equal(t, 281, d.Part2())
}
