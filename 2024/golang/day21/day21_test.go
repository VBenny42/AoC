package day21_test

import (
	"testing"

	"github.com/VBenny42/AoC/2024/golang/day21"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	day := day21.Parse("../../inputs/day21/sample-input.txt")

	assert.Equal(t, 126384, day.Part1())
}
