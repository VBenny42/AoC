package day22_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day22"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day22.Parse("../../inputs/day22/sample-input.txt")

	value, _ := day.Part1and2()

	assert.Equal(t, 37327623, value)
}

func TestPart2(t *testing.T) {
	day := day22.Parse("../../inputs/day22/sample-input.txt")

	_, value := day.Part1and2()

	assert.Equal(t, 23, value)
}
