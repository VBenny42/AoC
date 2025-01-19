package day22_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day22"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day22.Parse("../../inputs/day22/sample-input.txt", 4)

	assert.Equal(t, 6032, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day22.Parse("../../inputs/day22/sample-input.txt", 4)

	assert.Equal(t, 5031, d.Part2())
}
