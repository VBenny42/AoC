package day22_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day22"
)

func ExampleSolve() {
	day22.Solve("inputs/day22/sample-input.txt")
	// Output:
	// ANSWER1: winning player's score: 306
	// ANSWER2: winning player's score for recursive game: 291
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
