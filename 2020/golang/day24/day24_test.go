package day24_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day24"
)

func ExampleSolve() {
	day24.Solve("inputs/day24/sample-input.txt")
	// Output:
	// ANSWER1: number of tiles with black side up: 10
	// ANSWER2: number of black tiles after 100 days: 2208
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
		day24.Solve("inputs/day24/input.txt")
	}
}
