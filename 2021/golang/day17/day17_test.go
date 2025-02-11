package day17_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day17"
)

func ExampleSolve() {
	day17.Solve("inputs/day17/sample-input.txt")
	// Output:
	// ANSWER1: highest y-position that can be reached and still reach the target: 45
	// ANSWER2: number of distinct velocities that reach the target: 112
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
		day17.Solve("inputs/day17/input.txt")
	}
}
