package day19_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day19"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day19.Parse("inputs/day19/sample-input.txt")

	assert.Equal(t, 33, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day19.Parse("inputs/day19/sample-input.txt")

	assert.Equal(t, 3472, d.Part2())
}
