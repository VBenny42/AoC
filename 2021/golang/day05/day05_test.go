package day05_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day05"
)

func ExampleSolve() {
	day05.Solve("inputs/day05/sample-input.txt")
	// Output:
	// ANSWER1: number of points where at least two lines overlap: 5
	// ANSWER2: number of points where at least two lines overlap with diagonals: 12
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
		day05.Solve("inputs/day05/input.txt")
	}
}
