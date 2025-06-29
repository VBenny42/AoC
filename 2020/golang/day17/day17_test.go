package day17_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day17"
)

func ExampleSolve() {
	day17.Solve("inputs/day17/sample-input.txt")
	// Output:
	// ANSWER1: active cubes after 6 cycles: 112
	// ANSWER2: active cubes in 4D after 6 cycles: 848
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
