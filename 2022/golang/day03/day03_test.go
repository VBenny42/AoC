package day03_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day03"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day03.Parse("../../inputs/day03/sample-input.txt")

	assert.Equal(t, 157, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day03.Parse("../../inputs/day03/sample-input.txt")

	assert.Equal(t, 70, d.Part2())
}
