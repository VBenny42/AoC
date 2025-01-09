package day11_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day11"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day11.Parse("../../inputs/day11/sample-input.txt")

	assert.Equal(t, 10605, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day11.Parse("../../inputs/day11/sample-input.txt")

	assert.Equal(t, 2713310158, d.Part2())
}
