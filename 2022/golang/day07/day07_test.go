package day07_test

import (
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day07"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	d := day07.Parse("../../inputs/day07/sample-input.txt")

	assert.Equal(t, 48381165, d.Directories["/"])
	assert.Equal(t, 94853, d.Directories["/a/"])
	assert.Equal(t, 24933642, d.Directories["/d/"])
	assert.Equal(t, 584, d.Directories["/a/e/"])
}

func TestPart1(t *testing.T) {
	d := day07.Parse("../../inputs/day07/sample-input.txt")

	assert.Equal(t, 95437, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day07.Parse("../../inputs/day07/sample-input.txt")

	assert.Equal(t, 24933642, d.Part2())
}
