package day13_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day13"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day13.Parse("../../inputs/day13/sample-input.txt")

	assert.Equal(t, 13, day.Part1())
}

func TestPart2(t *testing.T) {
	day := day13.Parse("../../inputs/day13/sample-input.txt")

	assert.Equal(t, 140, day.Part2())
}
