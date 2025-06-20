package day12_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day12"
)

func ExampleSolve() {
	day12.Solve("inputs/day12/sample-input.txt")
	// Output:
	// ANSWER1: manhattan difference between ending and starting position: 25
	// ANSWER2: manhattan difference between ending and starting position with waypoint: 286
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
		day12.Solve("inputs/day12/input.txt")
	}
}
