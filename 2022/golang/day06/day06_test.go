package day06_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day06"
	"github.com/stretchr/testify/assert"
)

func TestPart1_1(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-1.txt")
	assert.Equal(t, 7, day.Part1())
}

func TestPart1_2(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-2.txt")
	assert.Equal(t, 5, day.Part1())
}

func TestPart1_3(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-3.txt")
	assert.Equal(t, 6, day.Part1())
}

func TestPart1_4(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-4.txt")
	assert.Equal(t, 10, day.Part1())
}

func TestPart1_5(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-5.txt")
	assert.Equal(t, 11, day.Part1())
}

func TestPart2_1(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-1.txt")
	assert.Equal(t, 19, day.Part2())
}

func TestPart2_2(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-2.txt")
	assert.Equal(t, 23, day.Part2())
}

func TestPart2_3(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-3.txt")
	assert.Equal(t, 23, day.Part2())
}

func TestPart2_4(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-4.txt")
	assert.Equal(t, 29, day.Part2())
}

func TestPart2_5(t *testing.T) {
	day := day06.Parse("inputs/day06/sample-input-5.txt")
	assert.Equal(t, 26, day.Part2())
}
