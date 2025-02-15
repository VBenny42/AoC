package day22_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day22"
)

func ExampleSolve() {
	day22.Solve("inputs/day22/sample-input.txt")
	// Output:
	// ANSWER1: number of cubes on: 590784
	// ANSWER2: number of cubes on using all steps: 39769202357779
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
		day22.Solve("inputs/day22/input.txt")
	}
}
