package day17_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day17"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day17.Parse("inputs/day17/sample-input.txt")

	assert.Equal(t, 3068, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day17.Parse("inputs/day17/sample-input.txt")

	assert.Equal(t, 1514285714288, d.Part2())
}

func BenchmarkSolve(b *testing.B) {
	b.StopTimer()

	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatalf("Failed to open %s: %s", os.DevNull, err)
	}
	defer devNull.Close()

	origStdout := os.Stdout
	os.Stdout = devNull

	b.StartTimer()

	for range b.N {
		day17.Solve("inputs/day17/sample-input.txt")
	}

	os.Stdout = origStdout
}
