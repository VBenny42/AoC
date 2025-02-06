package day04_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day04"
)

func ExampleSolve() {
	day04.Solve("inputs/day04/sample-input.txt")
	// Output:
	// ANSWER1: final score for first bingo board: 4512
	// ANSWER2: final score for last bingo board: 1924
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
		day04.Solve("inputs/day04/input.txt")
	}
}
