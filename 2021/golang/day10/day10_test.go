package day10_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day10"
)

func ExampleSolve() {
	day10.Solve("inputs/day10/sample-input.txt")
	// Output:
	// ANSWER1: total syntax error score: 26397
	// ANSWER2: total auto-correct score: 288957
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
		day10.Solve("inputs/day10/input.txt")
	}
}
