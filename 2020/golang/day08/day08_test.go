package day08_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day08"
)

func ExampleSolve() {
	day08.Solve("inputs/day08/sample-input.txt")
	// Output:
	// ANSWER1: accumuluator before infinite loop starts: 5
	// ANSWER2: accumuluator after fixing the program: 8
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
		day08.Solve("inputs/day08/input.txt")
	}
}
