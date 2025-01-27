package day11_test

import (
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day11"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day11.Parse("inputs/day11/sample-input.txt")

	val, _ := d.Part1And2()

	assert.Equal(t, 374, val)
}
