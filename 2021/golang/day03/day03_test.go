package day03_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day03"
)

func ExampleSolve() {
	day03.Solve("inputs/day03/sample-input.txt")
	// Output:
	// ANSWER1: power consumption of submarine: 198
	// ANSWER2: life support rating of submarine: 230
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
