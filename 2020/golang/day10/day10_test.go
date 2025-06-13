package day10_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day10"
)

func ExampleSolve() {
	day10.Solve("inputs/day10/sample-input.txt")
	day10.Solve("inputs/day10/sample-input1.txt")
	// Output:
	// ANSWER1: product of 1-jolt and 3-jolt differences: 35
	// ANSWER2: number of distinct arrangements: 8
	// ANSWER1: product of 1-jolt and 3-jolt differences: 220
	// ANSWER2: number of distinct arrangements: 19208
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
		day10.Solve("inputs/day10/input.txt")
	}
}
