package day17_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day17"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day17.Parse("../../inputs/day17/sample-input.txt")

	assert.Equal(t, 3068, d.Part1())
}
