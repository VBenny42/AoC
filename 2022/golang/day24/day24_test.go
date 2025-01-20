package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := Parse("../../inputs/day24/sample-input.txt")

	val, _ := d.Part1And2()

	assert.Equal(t, 18, val)
}

func TestPart2(t *testing.T) {
	d := Parse("../../inputs/day24/sample-input.txt")

	_, val := d.Part1And2()

	assert.Equal(t, 54, val)
}
