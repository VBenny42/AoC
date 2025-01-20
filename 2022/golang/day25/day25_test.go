package day25_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day25"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day25.Parse("inputs/day25/sample-input.txt")

	assert.Equal(t, "2=-1=0", d.Part1())
}
