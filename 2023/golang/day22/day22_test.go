package day22_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day22"
)

func ExampleSolve() {
	day22.Solve("inputs/day22/sample-input.txt")
	// Output:
	// ANSWER1: number of bricks that can be disintegrated without causing any other bricks to fall: 5
	// ANSWER2: number of bricks that will fall due to chain reaction: 7
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
