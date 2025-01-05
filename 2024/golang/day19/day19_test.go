package day19_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day19"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day19.Parse("../../inputs/day19/sample-input.txt")

	value, _ := day.Part1and2()

	assert.Equal(t, 6, value)
}

func TestPart2(t *testing.T) {
	day := day19.Parse("../../inputs/day19/sample-input.txt")

	_, value := day.Part1and2()

	assert.Equal(t, 16, value)
}
