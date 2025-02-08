package day09_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day09"
)

func ExampleSolve() {
	day09.Solve("inputs/day09/sample-input.txt")
	// Output:
	// ANSWER1: sum of risk levels of low points on map: 15
	// ANSWER2: product of the three largest basins: 1134
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
		day09.Solve("inputs/day09/input.txt")
	}
}
