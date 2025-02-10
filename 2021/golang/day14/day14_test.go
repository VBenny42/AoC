package day14_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day14"
)

func ExampleSolve() {
	day14.Solve("inputs/day14/sample-input.txt")
	// Output:
	// ANSWER1: most common element - least common element after 10 steps: 1588
	// ANSWER2: most common element - least common element after 40 steps: 2188189693529
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
