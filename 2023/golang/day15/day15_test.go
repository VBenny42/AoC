package day15_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day15"
)

func ExampleSolve() {
	day15.Solve("inputs/day15/sample-input.txt")
	// Output:
	// ANSWER1: sum of hashes: 1320
	// ANSWER2: focusing power of lens configuration: 145
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
		day15.Solve("inputs/day15/input.txt")
	}
}
