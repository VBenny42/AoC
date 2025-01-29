package day14_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day14"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day14.Parse("inputs/day14/sample-input.txt")

	assert.Equal(t, 136, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day14.Parse("inputs/day14/sample-input.txt")

	assert.Equal(t, 64, d.Part2())
}

func BenchmarkSolve(b *testing.B) {
	// b.ReportAllocs()

	oldStdout := os.Stdout
	null, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	os.Stdout = null

	defer func() {
		os.Stdout = oldStdout
		null.Close()
	}()

	b.ResetTimer()

	for range b.N {
		day14.Solve("inputs/day14/input.txt")
	}
}
