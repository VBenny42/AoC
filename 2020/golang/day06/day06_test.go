package day06_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day06"
)

func ExampleSolve() {
	day06.Solve("inputs/day06/sample-input.txt")
	// Output:
	// ANSWER1: sum of counts for `yes`: 11
	// ANSWER2: sum of counts for `yes` in all members: 6
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
		day06.Solve("inputs/day06/input.txt")
	}
}
