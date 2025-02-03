package day23_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day23"
)

func ExampleSolve() {
	day23.Solve("inputs/day23/sample-input.txt")
	// Output:
	// ANSWER1: longest path from start to end: 94
	// ANSWER2: longest path from start to end considering slopes as normal paths: 154
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
