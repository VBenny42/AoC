package day14_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day14"
	"github.com/VBenny42/AoC/2022/golang/utils"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	d := day14.Parse("../../inputs/day14/sample-input.txt")

	assert.Equal(t, len(d.Lines), 5)

	lines := []day14.Line{
		{Start: utils.Coord{X: 498, Y: 4}, End: utils.Coord{X: 498, Y: 6}},
		{Start: utils.Coord{X: 498, Y: 6}, End: utils.Coord{X: 496, Y: 6}},
		{Start: utils.Coord{X: 503, Y: 4}, End: utils.Coord{X: 502, Y: 4}},
		{Start: utils.Coord{X: 502, Y: 4}, End: utils.Coord{X: 502, Y: 9}},
		{Start: utils.Coord{X: 502, Y: 9}, End: utils.Coord{X: 494, Y: 9}},
	}

	assert.Equal(t, d.Lines, lines)
}

func TestPart1(t *testing.T) {
	d := day14.Parse("../../inputs/day14/sample-input.txt")

	value, _ := d.Part1And2()

	assert.Equal(t, 24, value)
}

func TestPart2(t *testing.T) {
	d := day14.Parse("../../inputs/day14/sample-input.txt")

	_, value := d.Part1And2()

	assert.Equal(t, 93, value)
}
