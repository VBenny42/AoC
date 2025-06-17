package day11_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day11"
)

func ExampleSolve() {
	day11.Solve("inputs/day11/sample-input.txt")
	// Output:
	// ANSWER1: number of occupied seats after no changes: 37
	// ANSWER2: number of occupied seats with visibility rules: 26
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
		day11.Solve("inputs/day11/input.txt")
	}
}
