package day19_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day19"
)

func ExampleSolve() {
	day19.Solve("inputs/day19/sample-input.txt")
	// Output:
	// ANSWER1: total beacons: 79
	// ANSWER2: largest manhattan distance between two scanners: 3621
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
		day19.Solve("inputs/day19/input.txt")
	}
}
