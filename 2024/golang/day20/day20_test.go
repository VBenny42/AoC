package day20_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day20"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day20.Parse("../../inputs/day20/sample-input.txt", 2)

	value, _ := day.Part1and2()

	assert.Equal(t, 44, value)
}

func TestPart2(t *testing.T) {
	day := day20.Parse("../../inputs/day20/sample-input.txt", 50)

	_, value := day.Part1and2()

	assert.Equal(t, 285, value)
}
