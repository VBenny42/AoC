package day02_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day02"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	d := day02.Parse("inputs/day02/sample-input.txt")

	want := []day02.Game{
		{Samples: []day02.Sample{
			{Blue: 3, Red: 4},
			{Red: 1, Green: 2, Blue: 6},
			{Green: 2},
		}},
		{Samples: []day02.Sample{
			{Blue: 1, Green: 2},
			{Green: 3, Blue: 4, Red: 1},
			{Green: 1, Blue: 1},
		}},
		{Samples: []day02.Sample{
			{Green: 8, Blue: 6, Red: 20},
			{Blue: 5, Red: 4, Green: 13},
			{Green: 5, Red: 1},
		}},
		{Samples: []day02.Sample{
			{Green: 1, Red: 3, Blue: 6},
			{Green: 3, Red: 6},
			{Green: 3, Blue: 15, Red: 14},
		}},
		{Samples: []day02.Sample{
			{Red: 6, Blue: 1, Green: 3},
			{Blue: 2, Red: 1, Green: 2},
		}},
	}

	assert.Equal(t, want, d.Games)
}

func TestPart1(t *testing.T) {
	d := day02.Parse("inputs/day02/sample-input.txt")

	assert.Equal(t, 8, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day02.Parse("inputs/day02/sample-input.txt")

	assert.Equal(t, 2286, d.Part2())
}
