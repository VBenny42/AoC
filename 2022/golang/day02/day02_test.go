package day02_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day02"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day02.Parse("inputs/day02/sample-input.txt")

	assert.Equal(t, 15, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day02.Parse("inputs/day02/sample-input.txt")

	assert.Equal(t, 12, d.Part2())
}
