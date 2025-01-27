package day10_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day10"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day10.Parse("inputs/day10/sample-input.txt")

	val, _ := d.Part1And2()

	assert.Equal(t, 4, val)
}

func TestPart1MoreComplex(t *testing.T) {
	d := day10.Parse("inputs/day10/sample-input-1.txt")

	val, _ := d.Part1And2()

	assert.Equal(t, 8, val)
}

func TestPart2(t *testing.T) {
	d := day10.Parse("inputs/day10/sample-input-2.txt")

	_, val := d.Part1And2()

	assert.Equal(t, 4, val)
}

func TestPart2MoreComplex(t *testing.T) {
	d := day10.Parse("inputs/day10/sample-input-3.txt")

	_, val := d.Part1And2()

	assert.Equal(t, 8, val)
}
