package day18_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day18"
)

func ExampleSolve() {
	day18.Solve("inputs/day18/sample-input.txt")
	// Output:
	// ANSWER1: magnitude of final sum: 4140
	// ANSWER2: maximum magnitude of any two different pairs: 3993
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
		day18.Solve("inputs/day18/input.txt")
	}
}
