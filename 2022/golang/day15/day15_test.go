package day15_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day15"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day15.Parse("../../inputs/day15/sample-input.txt", 10)

	assert.Equal(t, 26, day.Part1())
}