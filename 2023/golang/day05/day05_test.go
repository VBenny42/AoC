package day05_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day05"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day05.Parse("inputs/day05/sample-input.txt")

	assert.Equal(t, 35, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day05.Parse("inputs/day05/sample-input.txt")

	assert.Equal(t, 46, d.Part2())
}

func BenchmarkPart1(b *testing.B) {
	b.StopTimer()
	d := day05.Parse("inputs/day05/sample-input.txt")
	b.StartTimer()
	for range b.N {
		d.Part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	b.StopTimer()
	d := day05.Parse("inputs/day05/sample-input.txt")
	b.StartTimer()
	for range b.N {
		d.Part2()
	}
}
