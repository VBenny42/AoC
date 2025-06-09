package day03_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day03"
)

func ExampleSolve() {
	day03.Solve("inputs/day03/sample-input.txt")
	// Output:
	// ANSWER1: number of trees encountered: 7
	// ANSWER2: product of number of trees encountered on all slopes: 336
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
		day03.Solve("inputs/day03/input.txt")
	}
}
