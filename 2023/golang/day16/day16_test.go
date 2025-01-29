package day16_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day16"
)

func ExampleSolve() {
	day16.Solve("inputs/day16/sample-input.txt")
	// Output:
	// ANSWER1: number of energized tiles: 46
	// ANSWER2: max number of energized tiles possible: 51
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
		day16.Solve("inputs/day16/input.txt")
	}
}
