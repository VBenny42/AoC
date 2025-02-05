package day02_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day02"
)

func ExampleSolve() {
	day02.Solve("inputs/day02/sample-input.txt")
	// Output:
	// ANSWER1: product of final horizontal position and depth: 150
	// ANSWER2: product of final horizontal position and depth with aim: 900
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
		day02.Solve("inputs/day02/input.txt")
	}
}
