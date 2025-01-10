package day04_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day04"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day04.Parse("../../inputs/day04/sample-input.txt")

	assert.Equal(t, 2, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day04.Parse("../../inputs/day04/sample-input.txt")

	assert.Equal(t, 4, d.Part2())
}