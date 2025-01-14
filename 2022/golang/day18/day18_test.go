package day18_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day18"
	"github.com/stretchr/testify/assert"
)

func TestPart1_Smaller(t *testing.T) {
	d := day18.Parse("../../inputs/day18/sample-input-smaller.txt")

	assert.Equal(t, 10, d.Part1())
}

func TestPart1_Larger(t *testing.T) {
	d := day18.Parse("../../inputs/day18/sample-input-larger.txt")

	assert.Equal(t, 64, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day18.Parse("../../inputs/day18/sample-input-larger.txt")

	assert.Equal(t, 58, d.Part2())
}
