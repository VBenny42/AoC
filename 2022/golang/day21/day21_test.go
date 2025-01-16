package day21_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day21"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day21.Parse("../../inputs/day21/sample-input.txt")

	assert.Equal(t, 152, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day21.Parse("../../inputs/day21/sample-input.txt")

	assert.Equal(t, 301, d.Part2())
}
