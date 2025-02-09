package day11_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day11"
)

func ExampleSolve() {
	day11.Solve("inputs/day11/sample-input.txt")
	// Output:
	// ANSWER1: total flashes after 100 steps: 1656
	// ANSWER2: index where whole grid flashes: 195
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
