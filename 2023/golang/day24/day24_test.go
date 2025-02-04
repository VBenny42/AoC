package day24_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day24"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day24.Parse("inputs/day24/sample-input.txt", 7, 21)

	assert.Equal(t, 2, d.Part1())
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
		day24.Solve("inputs/day24/input.txt")
	}
}
