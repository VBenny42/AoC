package day18_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day18"
)

func ExampleSolve() {
	day18.Solve("inputs/day18/sample-input.txt")
	// Output:
	// ANSWER1: cubic metres of lava in trench: 62
	// ANSWER2: cubic metres of lava in trench after converting colors: 952408144115
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
