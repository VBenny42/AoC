package day10_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day10"
	"github.com/VBenny42/AoC/2022/golang/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day10.Parse("../../inputs/day10/sample-input.txt")

	assert.Equal(t, 13140, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day10.Parse("../../inputs/day10/sample-input.txt")

	file := utils.Must(os.ReadFile("../../inputs/day10/sample-output.txt"))

	output := string(file)

	assert.Equal(t, output, d.Part2())
}
