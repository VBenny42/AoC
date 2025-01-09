package day12_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day12"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day12.Parse("../../inputs/day12/sample-input.txt")

	value, _ := day.Part1And2()

	assert.Equal(t, 31, value)
}

func TestPart2(t *testing.T) {
	day := day12.Parse("../../inputs/day12/sample-input.txt")

	_, value := day.Part1And2()

	assert.Equal(t, 29, value)
}
