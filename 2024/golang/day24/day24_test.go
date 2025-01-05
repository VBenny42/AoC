package day24_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day24"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day24.Parse("../../inputs/day24/sample-input.txt")

	assert.Equal(t, int64(2024), day.Part1())
}
