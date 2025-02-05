package day01_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day01"
)

func ExampleSolve() {
	day01.Solve("inputs/day01/sample-input.txt")
	// Output:
	// ANSWER1: number of measurements larger than previous: 7
	// ANSWER2: number of measurements larger than previous for 3-measurement sliding window: 5
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
		day01.Solve("inputs/day01/input.txt")
	}
}
