package day23_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day23"
)

func ExampleSolve() {
	day23.Solve("inputs/day23/sample-input.txt")
	// Output:
	// ANSWER1: labels on the cups after cup `1`: 67384529
	// ANSWER2: product of the two cups after cup `1`: 149245887792
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
		day23.Solve("inputs/day23/input.txt")
	}
}
