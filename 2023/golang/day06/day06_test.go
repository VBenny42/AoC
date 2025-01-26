package day06_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day06"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day06.Parse("inputs/day06/sample-input.txt")

	assert.Equal(t, 288, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day06.Parse("inputs/day06/sample-input.txt")

	assert.Equal(t, 71503, d.Part2())
}

func BenchmarkPart1(b *testing.B) {
	d := day06.Parse("inputs/day06/sample-input.txt")

	for range b.N {
		d.Part1()
	}
}
