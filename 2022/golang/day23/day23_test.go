package day23_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day23"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day23.Parse("../../inputs/day23/sample-input.txt")

	assert.Equal(t, 110, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day23.Parse("../../inputs/day23/sample-input.txt")

	assert.Equal(t, 20, d.Part2())
}
