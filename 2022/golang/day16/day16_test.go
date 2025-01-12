package day16_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day16"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day16.Parse("../../inputs/day16/sample-input.txt")

	assert.Equal(t, 1651, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day16.Parse("../../inputs/day16/sample-input.txt")

	assert.Equal(t, 1707, d.Part2())
}
