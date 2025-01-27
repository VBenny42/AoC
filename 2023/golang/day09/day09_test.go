package day09_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day09"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day09.Parse("inputs/day09/sample-input.txt")

	assert.Equal(t, 114, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day09.Parse("inputs/day09/sample-input.txt")

	assert.Equal(t, 2, d.Part2())
}
